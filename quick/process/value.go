package process

import (
	"github.com/sohaha/zlsgo/ztime"
	"github.com/sohaha/zlsgo/ztype"
	"github.com/zlsgo/app_module/quick/define"
)

func (p *Process) InsertData(d *define.Define, data ztype.Map) (ztype.Map, error) {
	var err error
	data, err = p.ValuesBeforeProcess(data)
	if err != nil {
		return nil, err
	}

	data, err = p.VerifiData(data, ActiveCreate)
	if err != nil {
		return nil, err
	}

	if d.Options.Timestamps {
		data[define.Inside.CreatedAtKey()] = ztime.Time()
		data[define.Inside.UpdatedAtKey()] = ztime.Time()
	}

	if d.Options.SoftDeletes {
		data[define.Inside.DeletedAtKey()] = 0
	}

	return data, nil
}
