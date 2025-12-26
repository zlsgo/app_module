package model

import (
	"errors"
	"strconv"
	"time"

	"github.com/sohaha/zlsgo/zarray"
	"github.com/sohaha/zlsgo/zjson"
	"github.com/sohaha/zlsgo/ztime"
	"github.com/sohaha/zlsgo/ztype"
	mSchema "github.com/zlsgo/app_module/model/schema"
	"github.com/zlsgo/zdb/schema"
)

type activeType uint

const (
	activeCreate activeType = iota + 1
	activeUpdate
)

type ValidTriggerType uint

const (
	ValidTriggerAll ValidTriggerType = iota
	ValidTriggerCreate
	ValidTriggerUpdate
)

// filterDate 过滤数据字段
func filterDate(data ztype.Map, fields []string) ztype.Map {
	l := len(fields)
	if l == 0 {
		return data
	}

	n := make(ztype.Map, len(data))
	for k := range data {
		if !zarray.Contains(fields, k) {
			n[k] = data[k]
		}
	}

	return n
}

// VerifiData 验证数据
func VerifiData(data ztype.Map, columns mSchema.Fields, active activeType) (ztype.Map, error) {
	d := make(ztype.Map, len(columns))
	for name, column := range columns {
		if active == activeUpdate && column.Options.ReadOnly {
			continue
		}

		name, label := name, column.Label
		v, ok := data[name]

		{
			if !ok && active != activeUpdate {
				if column.Default != nil {
					v = column.Default
					if column.Type == schema.JSON {
						switch v.(type) {
						case string:
						default:
							v, _ = zjson.Marshal(v)
						}
					}
					ok = true
				}
			}

			if !ok && !column.Nullable && active != activeUpdate {
				return d, errors.New(label + "不能为空")
			}
		}

		if ok {
			if v == nil {
				if column.Nullable {
					continue
				} else {
					return d, errors.New(label + "不能为 null")
				}
			}
			if vt, ok := v.(ztype.Type); ok {
				v = vt.Value()
			}
			typ := column.Type
			switch typ {
			case schema.Bool:
				d[name] = ztype.ToBool(v)
	case schema.Time:
		err := column.GetValidations().VerifiAny(v).Error()
		if err != nil {
			return d, err
		}
		switch t := v.(type) {
		default:
			return d, errors.New(label + ": 未知时间格式")
		case DataTime:
			d[name] = t
		case time.Time:
			d[name] = DataTime{Time: t}
		case int, int8, int16, int32, int64:
			d[name] = DataTime{Time: ztime.Unix(ztype.ToInt64(v))}
		case uint, uint8, uint16, uint32, uint64:
			d[name] = DataTime{Time: ztime.Unix(ztype.ToInt64(v))}
		case float32, float64:
			d[name] = DataTime{Time: ztime.Unix(ztype.ToInt64(v))}
		case string:
			var (
				r   time.Time
				err error
			)
			// if column.Options.FormatTime == "" {
			r, err = ztime.Parse(t)
			// } else {
			// 	r, err = ztime.Parse(t, column.Options.FormatTime)
			// }
			if err != nil {
				if ts, err2 := strconv.ParseInt(t, 10, 64); err2 == nil {
					d[name] = DataTime{Time: ztime.Unix(ts)}
					break
				}
				return d, errors.New(label + ": 时间格式错误")
			}
			d[name] = DataTime{Time: r}
		}
	case schema.JSON:
		err := column.GetValidations().VerifiAny(v).Error()
		if err != nil {
					return d, err
				}
				d[name] = v
			default:
				var (
					val interface{}
					err error
				)

		switch typ {
		case schema.Bytes:
			b := ztype.ToBytes(v)
			val = b
			if column.Size > 0 && len(b) > int(column.Size) {
				return d, errors.New(label + "超过最大长度")
			}
				case schema.String, schema.Text:
					val, err = column.GetValidations().VerifiAny(v).String()
					if val == "" {
						if !column.Nullable {
							return d, errors.New(label + "不能为空")
						}
						val = ""
					}
				default:
					rule := column.GetValidations().VerifiAny(v)
					switch typ {
					case "int", "int8", "int16", "int32", "int64":
						val, err = rule.IsNumber().Int()
					case "uint", "uint8", "uint16", "uint32", "uint64":
						val, err = rule.IsNumber().Int()
						if err == nil {
							val = ztype.ToUint(val)
						}
					case "float", "float32", "float64":
						val, err = rule.IsNumber().Float64()
					default:
					}
				}

				if err != nil {
					return d, err
				}

				d[name] = val
			}
		}
	}

	return d, nil
}
