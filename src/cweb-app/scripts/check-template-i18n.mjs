import {readFile, readdir} from 'node:fs/promises';
import {join, resolve} from 'node:path';

const root = resolve('src/app/components');
const technicalTerms = /^(FreeSWITCH|CustomPBX|FS_CLI|CDR|HEP|Sofia|Verto|ACL|ESL|ODBC|XML|PNG|TXT|ASC|DESC|ORDER BY|on-true|on-false|always|never|true|false)$/i;
const visibleText = />([^<>{]*[A-Za-z][^<>{]*)</g;
const hardCodedAction = /<button\b[^>]*>\s*(Save|Add|Delete|Update|Remove|Submit|Cancel|Reset|Import|Copy)\s*<\/button>/gi;

// This is a debt ceiling, not an exemption list. Existing configuration templates
// still contain FreeSWITCH vocabulary and legacy labels; every migration lowers the
// ceiling, while new hard-coded UI copy makes CI fail immediately.
const baseline = {
  visibleText: 16,
  hardCodedActions: 0,
};

async function htmlFiles(directory) {
  const entries = await readdir(directory, {withFileTypes: true});
  const nested = await Promise.all(entries.map(entry => {
    const path = join(directory, entry.name);
    if (entry.isDirectory()) {
      return htmlFiles(path);
    }
    return entry.name.endsWith('.html') ? [path] : [];
  }));
  return nested.flat();
}

let visibleTextCount = 0;
let hardCodedActionCount = 0;
for (const file of await htmlFiles(root)) {
  const template = await readFile(file, 'utf8');
  const textMatches = [...template.matchAll(visibleText)]
    .map(match => match[1].trim())
    .filter(text => text && !text.includes('| transloco') && !technicalTerms.test(text));
  visibleTextCount += textMatches.length;
  hardCodedActionCount += [...template.matchAll(hardCodedAction)].length;
}

if (visibleTextCount > baseline.visibleText || hardCodedActionCount > baseline.hardCodedActions) {
  throw new Error(
    `Hard-coded UI text increased: visible ${visibleTextCount}/${baseline.visibleText}, ` +
    `actions ${hardCodedActionCount}/${baseline.hardCodedActions}. Use Transloco keys instead.`
  );
}

console.log(
  `Template i18n debt guard passed: ${visibleTextCount} visible strings, ` +
  `${hardCodedActionCount} hard-coded common actions (ceilings ${baseline.visibleText}/${baseline.hardCodedActions}).`
);
