package altStruct

type ConfigHttpCacheSetting struct {
	Id          int64               `xml:"-" json:"id" customsql:"pkey:id;check(id <> 0)"`
	Position    int64               `xml:"-" json:"position" customsql:"position;position"`
	Enabled     bool                `xml:"-" json:"enabled" customsql:"enabled;default=TRUE"`
	Name        string              `xml:"name,attr" json:"name" customsql:"param_name;unique;check(param_name <> '')"`
	Value       string              `xml:"value,attr" json:"value" customsql:"param_value"`
	Description string              `xml:"-" json:"description" customsql:"description"`
	Parent      *ConfigurationsList `xml:"-" json:"parent" customsql:"fkey:parent_id;unique;check(parent_id <> 0)"`
}

func (w *ConfigHttpCacheSetting) GetTableName() string {
	return "config_httpcache_settings"
}

type ConfigHttpCacheProfile struct {
	Id          int64               `xml:"-" json:"id" customsql:"pkey:id;check(id <> 0)"`
	Position    int64               `xml:"-" json:"position" customsql:"position;position"`
	Enabled     bool                `xml:"-" json:"enabled" customsql:"enabled;default=TRUE"`
	Name        string              `xml:"name,attr" json:"name" customsql:"name;unique;check(name <> '')"`
	Description string              `xml:"-" json:"description" customsql:"description"`
	Parent      *ConfigurationsList `xml:"-" json:"parent" customsql:"fkey:parent_id;unique;check(parent_id <> 0)"`
}

func (w *ConfigHttpCacheProfile) GetTableName() string {
	return "config_httpcache_profiles"
}

type ConfigHttpCacheProfileAWSS3 struct {
	Id              int64                   `xml:"-" json:"id" customsql:"pkey:id;check(id <> 0)"`
	Position        int64                   `xml:"-" json:"position" customsql:"position;position"`
	Enabled         bool                    `xml:"-" json:"enabled" customsql:"enabled;default=TRUE"`
	AccessKeyId     string                  `xml:"access-key-id" json:"access_key_id" customsql:"access_key_id;check(access_key_id <> '')"`
	SecretAccessKey string                  `xml:"secret-access-key" json:"secret_access_key" customsql:"secret_access_key;check(secret_access_key <> '')"`
	BaseDomain      string                  `xml:"base-domain" json:"base_domain" customsql:"base_domain"`
	Region          string                  `xml:"region" json:"region" customsql:"region;check(region <> '')"`
	Expires         int64                   `xml:"expires" json:"expires" customsql:"expires"`
	Description     string                  `xml:"-" json:"description" customsql:"description"`
	Parent          *ConfigHttpCacheProfile `xml:"-" json:"parent" customsql:"fkey:parent_id;unique;check(parent_id <> 0)"`
}

func (w *ConfigHttpCacheProfileAWSS3) GetTableName() string {
	return "config_httpcache_profile_awss3"
}

type ConfigHttpCacheProfileAzureBlob struct {
	Id              int64                   `xml:"-" json:"id" customsql:"pkey:id;check(id <> 0)"`
	Position        int64                   `xml:"-" json:"position" customsql:"position;position"`
	Enabled         bool                    `xml:"-" json:"enabled" customsql:"enabled;default=TRUE"`
	SecretAccessKey string                  `xml:"secret-access-key" json:"secret_access_key" customsql:"secret_access_key;check(secret_access_key <> '')"`
	Description     string                  `xml:"-" json:"description" customsql:"description"`
	Parent          *ConfigHttpCacheProfile `xml:"-" json:"parent" customsql:"fkey:parent_id;unique;check(parent_id <> 0)"`
}

func (w *ConfigHttpCacheProfileAzureBlob) GetTableName() string {
	return "config_httpcache_profile_azureblob"
}

type ConfigHttpCacheProfileDomain struct {
	Id          int64                   `xml:"-" json:"id" customsql:"pkey:id;check(id <> 0)"`
	Position    int64                   `xml:"-" json:"position" customsql:"position;position"`
	Enabled     bool                    `xml:"-" json:"enabled" customsql:"enabled;default=TRUE"`
	Name        string                  `xml:"name,attr" json:"name" customsql:"name;unique;check(name <> '')"`
	Description string                  `xml:"-" json:"description" customsql:"description"`
	Parent      *ConfigHttpCacheProfile `xml:"-" json:"parent" customsql:"fkey:parent_id;unique;check(parent_id <> 0)"`
}

func (w *ConfigHttpCacheProfileDomain) GetTableName() string {
	return "config_httpcache_profile_domains"
}
