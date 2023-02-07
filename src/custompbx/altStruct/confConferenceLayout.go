package altStruct

type ConfigConferenceLayout struct {
	Id             int64               `xml:"-" json:"id" customsql:"pkey:id;check(id <> 0)"`
	Position       int64               `xml:"-" json:"position" customsql:"position;position"`
	Enabled        bool                `xml:"-" json:"enabled" customsql:"enabled;default=TRUE"`
	Name           string              `xml:"name,attr" json:"name" customsql:"layout_name;unique;check(layout_name <> '')"`
	Auto3dPosition string              `xml:"auto-3d-position,attr" json:"auto-3d-position" customsql:"auto_3d_position"`
	Description    string              `xml:"-" json:"description" customsql:"description"`
	Parent         *ConfigurationsList `xml:"-" json:"parent" customsql:"fkey:parent_id;unique;check(parent_id <> 0)"`
}

func (w *ConfigConferenceLayout) GetTableName() string {
	return "config_conference_layouts"
}

type ConfigConferenceLayoutImage struct {
	Id            int64                   `xml:"-" json:"id" customsql:"pkey:id;check(id <> 0)"`
	Position      int64                   `xml:"-" json:"position" customsql:"position;position"`
	Enabled       bool                    `xml:"-" json:"enabled" customsql:"enabled;default=TRUE"`
	X             string                  `xml:"x,attr"  json:"x" customsql:"image_x"`
	Y             string                  `xml:"y,attr"  json:"y" customsql:"image_y"`
	Scale         string                  `xml:"scale,attr"  json:"scale" customsql:"image_scale"`
	Floor         string                  `xml:"floor,attr"  json:"floor" customsql:"image_floor"`
	FloorOnly     string                  `xml:"floor-only,attr"  json:"floor_only" customsql:"image_floor_only"`
	Hscale        string                  `xml:"hscale,attr"  json:"hscale" customsql:"image_hscale"`
	Overlap       string                  `xml:"overlap,attr"  json:"overlap" customsql:"image_overlap"`
	ReservationId string                  `xml:"reservation_id,attr"  json:"reservation_id" customsql:"image_reservation_id"`
	Zoom          string                  `xml:"zoom,attr"  json:"zoom" customsql:"image_zoom"`
	Description   string                  `xml:"-" json:"description" customsql:"description"`
	Parent        *ConfigConferenceLayout `xml:"-" json:"parent" customsql:"fkey:parent_id;unique;check(parent_id <> 0)"`
}

func (w *ConfigConferenceLayoutImage) GetTableName() string {
	return "config_conference_layout_images"
}

type ConfigConferenceLayoutGroup struct {
	Id          int64               `xml:"-" json:"id" customsql:"pkey:id;check(id <> 0)"`
	Position    int64               `xml:"-" json:"position" customsql:"position;position"`
	Enabled     bool                `xml:"-" json:"enabled" customsql:"enabled;default=TRUE"`
	Name        string              `xml:"name,attr" json:"name" customsql:"param_name;unique;check(param_name <> '')"`
	Description string              `xml:"-" json:"description" customsql:"description"`
	Parent      *ConfigurationsList `xml:"-" json:"parent" customsql:"fkey:parent_id;unique;check(parent_id <> 0)"`
}

func (w *ConfigConferenceLayoutGroup) GetTableName() string {
	return "config_conference_layouts_groups"
}

type ConfigConferenceLayoutGroupLayout struct {
	Id          int64                        `xml:"-" json:"id" customsql:"pkey:id;check(id <> 0)"`
	Position    int64                        `xml:"-" json:"position" customsql:"position;position"`
	Enabled     bool                         `xml:"-" json:"enabled" customsql:"enabled;default=TRUE"`
	Body        string                       `xml:",chardata"  json:"body" customsql:"layout_body;unique"`
	Description string                       `xml:"-" json:"description" customsql:"description"`
	Parent      *ConfigConferenceLayoutGroup `xml:"-" json:"parent" customsql:"fkey:parent_id;unique;check(parent_id <> 0)"`
}

func (w *ConfigConferenceLayoutGroupLayout) GetTableName() string {
	return "config_conference_layouts_group_layouts"
}
