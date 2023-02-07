package xmlcurl

import (
	"custompbx/altData"
	"custompbx/mainStruct"
	"custompbx/pbxcache"
	"encoding/xml"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const (
	sectionDirectory     = "directory"
	sectionConfiguration = "configuration"
	sectionDialplan      = "dialplan"
	continueTrue         = "true"
	breakOnFalse         = "on-false"
	breakOnTrue          = "on-true"
	breakNever           = "never"
	breakAlways          = "always"
	regexAny             = "any"
	regexAll             = "all"
	regexXor             = "xor"
	postPurpose          = "purpose"
	postGateways         = "gateways"
	postNetworkList      = "network-list"
	postUser             = "user"
	postDomain           = "domain"
	postGroup            = "group"
	postKeyValue         = "key_value"
	postTagName          = "tag_name"
	postCallerContext    = "Caller-Context"
	postProfile          = "profile"
	postCCQueue          = "CC-Queue"
	Header               = `<?xml version="1.0" encoding="UTF-8" standalone="no"?>` + "\n"
)

type DebugLog struct {
	Log []string
}

func (d *DebugLog) AddLogLine(line string) {
	d.Log = append(d.Log, line)
}

type Document struct {
	XMLName xml.Name  `xml:"document"`
	Type    string    `xml:"type,attr"`
	Section []Section `xml:"section"`
	Comment string    `xml:",comment"`
}

type Section struct {
	Name          string                    `xml:"name,attr"`
	Description   string                    `xml:"description,attr,omitempty"`
	Configuration *mainStruct.Configuration `xml:"configuration,omitempty"`
	Domain        []interface{}             `xml:"domain,omitempty"`
	Result        []Result                  `xml:"result,omitempty"`
	Context       *mainStruct.Context       `xml:"context,omitempty"`
}

type Result struct {
	Status string `xml:"status,attr"`
}

func NotFound(post map[string]string) []byte {
	document := newDocument()
	if post != nil {
		document.Comment = fmt.Sprintf("section: %s, keyValue: %s, purpose: %s, user: %s, domain: %s, profile: %s, context: %s, tagName: %s.",
			post["section"], post[postKeyValue], post[postPurpose], post[postUser], post[postDomain], post[postProfile], post[postCallerContext], post[postTagName])
	}
	section := newSection("result")
	section.addResult("not found")
	document.addSection(section)

	res, ok := XmlToByte(document)
	if !ok {
		return []byte{}
	}
	return res
}

func Directory(post map[string]string) []byte {
	document := newDocument()
	section := newSection(sectionDirectory)
	section.Description = "FreeSWITCH Directory"

	/*	document.Comment = fmt.Sprintf("section: %s, keyValue: %s, purpose: %s, use: %s, domain: %s, profile: %s, context: %s",
		post["section"], post[postKeyValue], post[postPurpose], post[postUser], post[postDomain], post[postProfile], post[postCallerContext])
	*/
	var domains []interface{}
	switch true {
	case post[postPurpose] == postGateways:
		// profileName, ok := post[postProfile] //?????
		domains = altData.XMLDomainDirectoryGateways()
	case post[postPurpose] == postNetworkList:
		if post[postKeyValue] == "" || post[postTagName] != postDomain {
			return NotFound(post)
		}
		domains = altData.XMLDomainDirectoryNetworkLists(post[postKeyValue])
	case post[postUser] != "":
		if post[postDomain] == "" {
			return NotFound(post)
		}
		domains = altData.XMLDomainDirectoryUser(post[postDomain], post[postUser])
	case post[postGroup] != "":
		domainName, ok := post[postDomain]
		if !ok {
			return NotFound(post)
		}
		domains = altData.XMLDomainDirectoryUserGroup(domainName, post[postGroup])
	default:
		domains = altData.XMLDomainDirectoryDefault()
	}

	if len(domains) == 0 {
		return NotFound(post)
	}
	section.Domain = domains
	document.addSection(section)
	return prepareDocument(document)
}

func Configuration(post map[string]string) []byte {
	var ok bool

	document := newDocument()
	section := newSection(sectionConfiguration)
	section.Description = "FreeSWITCH Configuration"

	configName, ok := post[postKeyValue]

	if ok {
		config := altData.GetConfigSection(configName, post)
		if config == nil {
			return NotFound(post)
		}

		section.Configuration = config
		document.addSection(section)
		document.Comment = fmt.Sprintf("section: %s, keyValue: %s, purpose: %s, user: %s, domain: %s, profile: %s, context: %s",
			post["section"], post[postKeyValue], post[postPurpose], post[postUser], post[postDomain], post[postProfile], post[postCallerContext])
		return prepareDocument(document)
	}
	return NotFound(post)
}

func Dialplan(post map[string]string, events chan interface{}) []byte {
	var ok bool

	document := newDocument()
	section := newSection(sectionDialplan)
	section.Description = "FreeSWITCH Dialplan"

	contextName, ok := post[postCallerContext]
	if !ok {
		return NotFound(post)
	}
	context := pbxcache.GetContextByName(contextName)
	if context == nil {
		return NotFound(post)
	}

	var actions []mainStruct.Action
	extName := ""
	debugLog := DebugLog{}

	r := regexp.MustCompile(`^eavesdrop::(.+)$`)
	res := r.FindStringSubmatch(post["Caller-Destination-Number"])
	if len(res) > 1 && res[1] != "" {
		actions = append(actions, mainStruct.Action{Application: "answer", Data: ""})
		actions = append(actions, mainStruct.Action{Application: "eavesdrop", Data: res[1]})
	} else if context.Dialplan.NoProceed {
		return context.FullXMLCache.Get()
	} else {
		actions, extName = huntActions(post, context, &debugLog)
		if pbxcache.GetDialplanDebug() {
			events <- &mainStruct.DialplanDebug{Log: debugLog.Log, Actions: actions}
		}
	}

	section.Context = &mainStruct.Context{
		Name: contextName,
		XMLExtension: mainStruct.Extension{
			Name: extName,
			XMLCondition: mainStruct.Condition{
				XMLAction: actions,
			},
		},
	}

	document.addSection(section)
	return prepareDocument(document)
}

func prepareDocument(document Document) []byte {
	res, ok := XmlToByte(document)
	if !ok {
		return NotFound(nil)
	}
	return res
}

func XmlToByte(value ...interface{}) ([]byte, bool) {
	output, err := xml.MarshalIndent(value, "", "  ")
	if err != nil {
		log.Printf("%+v", err)
		return []byte{}, false
	}
	output = append(output, '\n')
	return output, true
}

func newDocument() Document {
	return Document{Type: "freeswitch/xml"}
}

func newSection(name string) Section {
	return Section{Name: name}
}

func (document *Document) addSection(section Section) {
	document.Section = append(document.Section, section)
}

func (section *Section) addResult(status string) {
	section.Result = append(section.Result, Result{
		Status: status,
	})
}

func huntActions(post map[string]string, context *mainStruct.Context, debugLog *DebugLog) ([]mainStruct.Action, string) {
	extensionName := "ext"
	extensions := context.Extensions.ExtensionsList()
	//log.Printf("NUMBER OF extensions IS %d\n", len(extensions))

	var result []mainStruct.Action
	//change range to i++
	for extIndex := 0; extIndex < len(extensions); extIndex++ {
		stackedCondition := true
		conditions := extensions[extIndex].Conditions.ConditionsList()
		for condIndex := 0; condIndex < len(conditions); condIndex++ {
			antiActions := conditions[condIndex].AntiActions.ActionsList()
			actions := conditions[condIndex].Actions.ActionsList()
			if !stackedCondition {
				if len(antiActions) != 0 || len(actions) != 0 {
					stackedCondition = true
				}
				continue
			}
			var ok bool
			var matchedGroups []string
			if ok, matchedGroups = checkCondition(conditions[condIndex], post, debugLog); !ok {
				if len(antiActions) == 0 {
					if len(actions) == 0 {
						stackedCondition = false
					}
					continue
				}
				stackedCondition = true
				extensionName += "_" + extensions[extIndex].Name
				for aActIndex := 0; aActIndex < len(antiActions); aActIndex++ {
					result = append(result, mainStruct.Action{Application: antiActions[aActIndex].Application, Data: antiActions[aActIndex].Data})
				}
				if extensions[extIndex].Continue != continueTrue {
					return result, extensionName
				}
				if conditions[condIndex].Break == "" || conditions[condIndex].Break == breakOnFalse || conditions[condIndex].Break == breakAlways {
					break
				}
				continue
			}
			stackedCondition = true
			for actIndex := 0; actIndex < len(actions); actIndex++ {
				//log.Printf("%+v - %+v - %+v - %+v", cond.Id, cond.Position, action.Id, action.Position)
				actionData := actions[actIndex].Data
				if len(matchedGroups) > 1 {
					for i := 1; i < len(matchedGroups); i++ {
						actionData = strings.Replace(actionData, "$"+strconv.Itoa(i), matchedGroups[i], -1)
					}
				}
				//HZ HZ changing original dialplan?
				if actions[actIndex].Inline && actions[actIndex].Application == "set" {
					data := strings.Split(actionData, "=")
					if len(data) > 1 {
						// conditions[condIndex].Field = "variable_" + data[0]
						// conditions[condIndex].Expression = data[1]
						post["variable_"+data[0]] = data[1]
					}
				}
				result = append(result, mainStruct.Action{Application: actions[actIndex].Application, Data: actionData, Inline: actions[actIndex].Inline})
			}
			if extensions[extIndex].Continue != continueTrue {
				return result, extensionName + "_" + extensions[extIndex].Name
			}
			extensionName += "_" + extensions[extIndex].Name

			if conditions[condIndex].Break == breakOnTrue || conditions[condIndex].Break == breakAlways {
				break
			}
		}
	}

	return result, extensionName
}

func checkCondition(condition *mainStruct.Condition, post map[string]string, debugLog *DebugLog) (bool, []string) {

	switch true {
	case condition.Field != "":
		return checkExpressionCondition(condition, post, debugLog)
	case condition.Regex != "":
		regexList := condition.Regexes.RegexList()
		return checkRegexes(regexList, post, condition.Regex, debugLog), []string{}
	case condition.Year != "":
		year := time.Now().Year()
		_, ok := matchList(condition.Year)[year]

		if !ok {
			return false, []string{}
		}
		condition.Year = ""
		return checkCondition(condition, post, debugLog)
	case condition.Mon != "":
		month := int(time.Now().Month())
		_, ok := matchList(condition.Year)[month]
		if !ok {
			return false, []string{}
		}
		condition.Mon = ""
		return checkCondition(condition, post, debugLog)
	case condition.Week != "":
		_, week := time.Now().ISOWeek()
		_, ok := matchList(condition.Year)[week]
		if !ok {
			return false, []string{}
		}
		condition.Week = ""
		return checkCondition(condition, post, debugLog)
	case condition.Wday != "":
		weekday := int(time.Now().Weekday()) + 1
		_, ok := matchList(condition.Year)[weekday]
		if !ok {
			return false, []string{}
		}
		condition.Wday = ""
		return checkCondition(condition, post, debugLog)
	case condition.Hour != "":
		hour := time.Now().Hour()
		_, ok := matchList(condition.Year)[hour]
		if !ok {
			return false, []string{}
		}
		condition.Hour = ""
		return checkCondition(condition, post, debugLog)
	case condition.Minute != "":
		minute := time.Now().Minute()
		_, ok := matchList(condition.Year)[minute]
		if !ok {
			return false, []string{}
		}
		condition.Minute = ""
		return checkCondition(condition, post, debugLog)
	case condition.Mday != "":
		day := time.Now().Day()
		_, ok := matchList(condition.Year)[day]
		if !ok {
			return false, []string{}
		}
		condition.Mday = ""
		return checkCondition(condition, post, debugLog)
	case condition.Yday != "":
		day := time.Now().YearDay()
		_, ok := matchList(condition.Year)[day]
		if !ok {
			return false, []string{}
		}
		condition.Yday = ""
		return checkCondition(condition, post, debugLog)
	case condition.DateTime != "":
		if !matchListDate(condition.TimeOfDay) {
			return false, []string{}
		}
		condition.DateTime = ""
		return checkCondition(condition, post, debugLog)
	case condition.TimeOfDay != "":
		if !matchListTime(condition.TimeOfDay) {
			return false, []string{}
		}
		condition.TimeOfDay = ""
		return checkCondition(condition, post, debugLog)

	// case condition.Mweek != "":
	// case condition.Minday != "":
	// case condition.Dst != "":
	// case condition.TzOffset != "":
	default:
		LogDialplan(true, condition, "", debugLog)
		return true, []string{}
	}
}

func normalizeField(field string, post map[string]string) []string {
	//global variables now only when have internally
	r := regexp.MustCompile(`\$\${(.+)}`)
	res := r.FindStringSubmatch(field)
	if len(res) > 1 {
		gvar := pbxcache.GetGlobalVariableByName(res[1])
		if gvar == nil {
			return []string{}
		}
		return []string{gvar.Value}
	}

	switch field {
	case "caller_id_name":
		return []string{post["Caller-Caller-ID-Name"]}
	case "caller_id_number":
		return []string{post["Caller-Caller-ID-Number"]}
	case "destination_number":
		return []string{post["Caller-Destination-Number"]}
	case "direction":
		return []string{post["Call-Direction"]}
	case "channel_name":
		return []string{post["Channel-Name"]}
	case "state":
		return []string{post["Channel-State"]}
	case "call_timeout":
	case "bridge_hangup_cause":
	}

	r = regexp.MustCompile(`\${(.+?)(?:\((.*)\))?}`) //TODO bracket can be replaced with space
	res = r.FindStringSubmatch(field)
	switch len(res) {
	case 2:
		return []string{"variable_" + res[1]}
	case 3:
		return []string{"", res[1], res[2]}
	}
	return []string{}
}

func getField(field string, post map[string]string) string {
	var command string
	var attr string

	fieldVal := normalizeField(field, post)
	switch len(fieldVal) {
	case 0:
		return ""
	case 1:
		return fieldVal[0]
	case 3:
		command = fieldVal[1]
		attr = fieldVal[2]
	}
	//only cond app for now
	var result string
	switch command {
	case "cond":
		r := regexp.MustCompile(`^\s*?([^\s]+?)\s*?([><=!][><=!]?)\s*?([^\s]+?)\s*?\?\s*?([^\s]+?)\s*?:\s*?([^\s]+)\s*?$`)
		res := r.FindStringSubmatch(attr)
		if len(res) != 6 {
			return ""
		}
		var subres bool
		fieldValues := normalizeField(res[1], post)
		if len(fieldValues) > 0 && fieldValues[0] != "" {
			res[1] = fieldValues[0]
		}
		switch res[2] {
		case "==":
			subres = res[1] == res[3]
		case ">":
			subres = res[1] > res[3]
		case ">=":
			subres = res[1] >= res[3]
		case "<":
			subres = res[1] < res[3]
		case "<=":
			subres = res[1] <= res[3]
		case "!=":
			subres = res[1] != res[3]
		}
		result = res[4]
		if !subres {
			result = res[5]
		}
	}
	return result
}

func checkField(expression, value string) (bool, []string) {
	r, err := regexp.Compile(expression)
	if err != nil {
		log.Printf("[CONDITION] [ERROR] %+v \n", err)
		return false, []string{}
	}

	match := r.FindStringSubmatch(value)

	matched := len(match) > 0

	return matched, match
}

func checkExpressionCondition(condition *mainStruct.Condition, post map[string]string, debugLog *DebugLog) (bool, []string) {
	field := getField(condition.Field, post)
	if field == "" {
		LogDialplan(false, condition, field, debugLog)
		return false, []string{}
	}

	match, matchedGroups := checkField(condition.Expression, field)

	LogDialplan(match, condition, field, debugLog)

	return match, matchedGroups
}

func checkRegexExpressionCondition(regex *mainStruct.Regex, post map[string]string, debugLog *DebugLog) bool {
	field := getField(regex.Field, post)
	if field == "" {
		LogDialplanRegex(false, regex, field, debugLog)
		return false
	}

	match, _ := checkField(regex.Expression, field)

	LogDialplanRegex(match, regex, field, debugLog)

	return match
}

func matchList(condition string) map[int]bool {
	result := map[int]bool{}
	parts := strings.Split(condition, ",")
	for _, v := range parts {
		match := strings.Contains(v, "-")
		if match {
			parts2 := strings.Split(condition, "-")
			if len(parts2) != 2 {
				break
			}
			int1, err := strconv.Atoi(parts2[0])
			if err != nil {
				break
			}
			int2, err := strconv.Atoi(parts2[1])
			if err != nil {
				break
			}
			for i := int2; int1 <= i; i-- {
				result[i] = true
			}
		} else {
			res, err := strconv.Atoi(v)
			if err != nil {
				break
			}
			result[res] = true
		}
	}
	return result
}

func matchListTime(condition string) bool {
	parts := strings.Split(condition, "-")

	if len(parts) != 2 {
		return false
	}
	timeEnd, err := time.Parse("03:04:05", parts[1])
	if err != nil {
		return false
	}
	timeStart, err := time.Parse("03:04:05", parts[0])
	if err != nil {
		return false
	}
	now := time.Now()

	res1 := timeEnd.Sub(now)
	if res1.Nanoseconds() > 0 {
		return false
	}

	res2 := timeStart.Sub(now)
	if res2.Nanoseconds() < 0 {
		return false
	}

	return true
}

func matchListDate(condition string) bool {
	parts := strings.Split(condition, "~")

	if len(parts) != 2 {
		return false
	}
	timeEnd, err := time.Parse(time.RFC3339, parts[1])
	if err != nil {
		return false
	}
	timeStart, err := time.Parse(time.RFC3339, parts[0])
	if err != nil {
		return false
	}
	now := time.Now()

	res1 := timeEnd.Sub(now)
	if res1.Nanoseconds() > 0 {
		return false
	}

	res2 := timeStart.Sub(now)
	if res2.Nanoseconds() < 0 {
		return false
	}

	return true
}

func LogDialplan(match bool, condition *mainStruct.Condition, expression string, debugLog *DebugLog) {
	if !pbxcache.GetDialplanDebug() {
		return
	}
	if expression == "" {
		expression = "_UNDEF_"
	}
	breakCondition := breakOnFalse
	if condition.Break != "" {
		breakCondition = condition.Break
	}
	fieldCondition := "EMPTY"
	if condition.Field != "" {
		fieldCondition = condition.Field
	}
	debugLog.AddLogLine(
		fmt.Sprintf("[CONDITION %d] (%t) [%+v] %+v(%+v) =~ /%+v/ break=%+v \n",
			condition.Id, match, condition.Extension.Name, fieldCondition, expression, condition.Expression, breakCondition),
	)
}

func LogDialplanRegex(match bool, regex *mainStruct.Regex, expression string, debugLog *DebugLog) {
	if !pbxcache.GetDialplanDebug() {
		return
	}
	if expression == "" {
		expression = "_UNDEF_"
	}

	fieldCondition := "EMPTY"
	if regex.Field != "" {
		fieldCondition = regex.Field
	}
	debugLog.AddLogLine(
		fmt.Sprintf("[CONDITION %d] [REGEX=%+v] (%t) [%+v] %+v(%+v) =~ /%+v/ \n",
			regex.Condition.Id, regex.Condition.Regex, match, regex.Condition.Extension.Name, fieldCondition, expression, regex.Expression),
	)
}

func checkRegexes(regexList []*mainStruct.Regex, post map[string]string, regCase string, debugLog *DebugLog) bool {
	var result bool
	var xorFlag bool
	switch regCase {
	case regexAny:
	case regexAll:
	case regexXor:
	default:
		return false
	}

	for _, reg := range regexList {
		result = checkRegexExpressionCondition(reg, post, debugLog)
		if (result && regCase == regexAny) || (!result && regCase == regexAll) {
			return result
		}
		if result && regCase == regexXor {
			if xorFlag {
				return !result
			}
			xorFlag = true
		}
	}

	if xorFlag {
		return xorFlag
	}

	return result
}
