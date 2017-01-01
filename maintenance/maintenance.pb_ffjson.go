// DO NOT EDIT!
// Code generated by ffjson <https://github.com/pquerna/ffjson>
// source: maintenance/maintenance.pb.go
// DO NOT EDIT!

package maintenance

import (
	"bytes"
	"github.com/mesos/mesos-go"
	"github.com/mesos/mesos-go/allocator"
	"fmt"
	fflib "github.com/pquerna/ffjson/fflib/v1"
)

func (mj *ClusterStatus) MarshalJSON() ([]byte, error) {
	var buf fflib.Buffer
	if mj == nil {
		buf.WriteString("null")
		return buf.Bytes(), nil
	}
	err := mj.MarshalJSONBuf(&buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
func (mj *ClusterStatus) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	if mj == nil {
		buf.WriteString("null")
		return nil
	}
	var err error
	var obj []byte
	_ = obj
	_ = err
	buf.WriteString(`{ `)
	if len(mj.DrainingMachines) != 0 {
		buf.WriteString(`"draining_machines":`)
		if mj.DrainingMachines != nil {
			buf.WriteString(`[`)
			for i, v := range mj.DrainingMachines {
				if i != 0 {
					buf.WriteString(`,`)
				}

				{

					if v == nil {
						buf.WriteString("null")
						return nil
					}

					err = v.MarshalJSONBuf(buf)
					if err != nil {
						return err
					}

				}
			}
			buf.WriteString(`]`)
		} else {
			buf.WriteString(`null`)
		}
		buf.WriteByte(',')
	}
	if len(mj.DownMachines) != 0 {
		buf.WriteString(`"down_machines":`)
		if mj.DownMachines != nil {
			buf.WriteString(`[`)
			for i, v := range mj.DownMachines {
				if i != 0 {
					buf.WriteString(`,`)
				}

				{

					if v == nil {
						buf.WriteString("null")
						return nil
					}

					err = v.MarshalJSONBuf(buf)
					if err != nil {
						return err
					}

				}
			}
			buf.WriteString(`]`)
		} else {
			buf.WriteString(`null`)
		}
		buf.WriteByte(',')
	}
	buf.Rewind(1)
	buf.WriteByte('}')
	return nil
}

const (
	ffj_t_ClusterStatusbase = iota
	ffj_t_ClusterStatusno_such_key

	ffj_t_ClusterStatus_DrainingMachines

	ffj_t_ClusterStatus_DownMachines
)

var ffj_key_ClusterStatus_DrainingMachines = []byte("draining_machines")

var ffj_key_ClusterStatus_DownMachines = []byte("down_machines")

func (uj *ClusterStatus) UnmarshalJSON(input []byte) error {
	fs := fflib.NewFFLexer(input)
	return uj.UnmarshalJSONFFLexer(fs, fflib.FFParse_map_start)
}

func (uj *ClusterStatus) UnmarshalJSONFFLexer(fs *fflib.FFLexer, state fflib.FFParseState) error {
	var err error = nil
	currentKey := ffj_t_ClusterStatusbase
	_ = currentKey
	tok := fflib.FFTok_init
	wantedTok := fflib.FFTok_init

mainparse:
	for {
		tok = fs.Scan()
		//	println(fmt.Sprintf("debug: tok: %v  state: %v", tok, state))
		if tok == fflib.FFTok_error {
			goto tokerror
		}

		switch state {

		case fflib.FFParse_map_start:
			if tok != fflib.FFTok_left_bracket {
				wantedTok = fflib.FFTok_left_bracket
				goto wrongtokenerror
			}
			state = fflib.FFParse_want_key
			continue

		case fflib.FFParse_after_value:
			if tok == fflib.FFTok_comma {
				state = fflib.FFParse_want_key
			} else if tok == fflib.FFTok_right_bracket {
				goto done
			} else {
				wantedTok = fflib.FFTok_comma
				goto wrongtokenerror
			}

		case fflib.FFParse_want_key:
			// json {} ended. goto exit. woo.
			if tok == fflib.FFTok_right_bracket {
				goto done
			}
			if tok != fflib.FFTok_string {
				wantedTok = fflib.FFTok_string
				goto wrongtokenerror
			}

			kn := fs.Output.Bytes()
			if len(kn) <= 0 {
				// "" case. hrm.
				currentKey = ffj_t_ClusterStatusno_such_key
				state = fflib.FFParse_want_colon
				goto mainparse
			} else {
				switch kn[0] {

				case 'd':

					if bytes.Equal(ffj_key_ClusterStatus_DrainingMachines, kn) {
						currentKey = ffj_t_ClusterStatus_DrainingMachines
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffj_key_ClusterStatus_DownMachines, kn) {
						currentKey = ffj_t_ClusterStatus_DownMachines
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				}

				if fflib.EqualFoldRight(ffj_key_ClusterStatus_DownMachines, kn) {
					currentKey = ffj_t_ClusterStatus_DownMachines
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffj_key_ClusterStatus_DrainingMachines, kn) {
					currentKey = ffj_t_ClusterStatus_DrainingMachines
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				currentKey = ffj_t_ClusterStatusno_such_key
				state = fflib.FFParse_want_colon
				goto mainparse
			}

		case fflib.FFParse_want_colon:
			if tok != fflib.FFTok_colon {
				wantedTok = fflib.FFTok_colon
				goto wrongtokenerror
			}
			state = fflib.FFParse_want_value
			continue
		case fflib.FFParse_want_value:

			if tok == fflib.FFTok_left_brace || tok == fflib.FFTok_left_bracket || tok == fflib.FFTok_integer || tok == fflib.FFTok_double || tok == fflib.FFTok_string || tok == fflib.FFTok_bool || tok == fflib.FFTok_null {
				switch currentKey {

				case ffj_t_ClusterStatus_DrainingMachines:
					goto handle_DrainingMachines

				case ffj_t_ClusterStatus_DownMachines:
					goto handle_DownMachines

				case ffj_t_ClusterStatusno_such_key:
					err = fs.SkipField(tok)
					if err != nil {
						return fs.WrapErr(err)
					}
					state = fflib.FFParse_after_value
					goto mainparse
				}
			} else {
				goto wantedvalue
			}
		}
	}

handle_DrainingMachines:

	/* handler: uj.DrainingMachines type=[]*maintenance.ClusterStatus_DrainingMachine kind=slice quoted=false*/

	{

		{
			if tok != fflib.FFTok_left_brace && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for ", tok))
			}
		}

		if tok == fflib.FFTok_null {
			uj.DrainingMachines = nil
		} else {

			uj.DrainingMachines = []*ClusterStatus_DrainingMachine{}

			wantVal := true

			for {

				var tmp_uj__DrainingMachines *ClusterStatus_DrainingMachine

				tok = fs.Scan()
				if tok == fflib.FFTok_error {
					goto tokerror
				}
				if tok == fflib.FFTok_right_brace {
					break
				}

				if tok == fflib.FFTok_comma {
					if wantVal == true {
						// TODO(pquerna): this isn't an ideal error message, this handles
						// things like [,,,] as an array value.
						return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
					}
					continue
				} else {
					wantVal = true
				}

				/* handler: tmp_uj__DrainingMachines type=*maintenance.ClusterStatus_DrainingMachine kind=ptr quoted=false*/

				{
					if tok == fflib.FFTok_null {

						tmp_uj__DrainingMachines = nil

						state = fflib.FFParse_after_value
						goto mainparse
					}

					if tmp_uj__DrainingMachines == nil {
						tmp_uj__DrainingMachines = new(ClusterStatus_DrainingMachine)
					}

					err = tmp_uj__DrainingMachines.UnmarshalJSONFFLexer(fs, fflib.FFParse_want_key)
					if err != nil {
						return err
					}
					state = fflib.FFParse_after_value
				}

				uj.DrainingMachines = append(uj.DrainingMachines, tmp_uj__DrainingMachines)

				wantVal = false
			}
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_DownMachines:

	/* handler: uj.DownMachines type=[]*mesos.MachineID kind=slice quoted=false*/

	{

		{
			if tok != fflib.FFTok_left_brace && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for ", tok))
			}
		}

		if tok == fflib.FFTok_null {
			uj.DownMachines = nil
		} else {

			uj.DownMachines = []*mesos.MachineID{}

			wantVal := true

			for {

				var tmp_uj__DownMachines *mesos.MachineID

				tok = fs.Scan()
				if tok == fflib.FFTok_error {
					goto tokerror
				}
				if tok == fflib.FFTok_right_brace {
					break
				}

				if tok == fflib.FFTok_comma {
					if wantVal == true {
						// TODO(pquerna): this isn't an ideal error message, this handles
						// things like [,,,] as an array value.
						return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
					}
					continue
				} else {
					wantVal = true
				}

				/* handler: tmp_uj__DownMachines type=*mesos.MachineID kind=ptr quoted=false*/

				{
					if tok == fflib.FFTok_null {

						tmp_uj__DownMachines = nil

						state = fflib.FFParse_after_value
						goto mainparse
					}

					if tmp_uj__DownMachines == nil {
						tmp_uj__DownMachines = new(mesos.MachineID)
					}

					err = tmp_uj__DownMachines.UnmarshalJSONFFLexer(fs, fflib.FFParse_want_key)
					if err != nil {
						return err
					}
					state = fflib.FFParse_after_value
				}

				uj.DownMachines = append(uj.DownMachines, tmp_uj__DownMachines)

				wantVal = false
			}
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

wantedvalue:
	return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
wrongtokenerror:
	return fs.WrapErr(fmt.Errorf("ffjson: wanted token: %v, but got token: %v output=%s", wantedTok, tok, fs.Output.String()))
tokerror:
	if fs.BigError != nil {
		return fs.WrapErr(fs.BigError)
	}
	err = fs.Error.ToError()
	if err != nil {
		return fs.WrapErr(err)
	}
	panic("ffjson-generated: unreachable, please report bug.")
done:
	return nil
}

func (mj *ClusterStatus_DrainingMachine) MarshalJSON() ([]byte, error) {
	var buf fflib.Buffer
	if mj == nil {
		buf.WriteString("null")
		return buf.Bytes(), nil
	}
	err := mj.MarshalJSONBuf(&buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
func (mj *ClusterStatus_DrainingMachine) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	if mj == nil {
		buf.WriteString("null")
		return nil
	}
	var err error
	var obj []byte
	_ = obj
	_ = err
	buf.WriteString(`{ `)
	if mj.Id != nil {
		if true {
			buf.WriteString(`"id":`)

			{

				err = mj.Id.MarshalJSONBuf(buf)
				if err != nil {
					return err
				}

			}
			buf.WriteByte(',')
		}
	}
	if len(mj.Statuses) != 0 {
		buf.WriteString(`"statuses":`)
		if mj.Statuses != nil {
			buf.WriteString(`[`)
			for i, v := range mj.Statuses {
				if i != 0 {
					buf.WriteString(`,`)
				}

				{

					if v == nil {
						buf.WriteString("null")
						return nil
					}

					err = v.MarshalJSONBuf(buf)
					if err != nil {
						return err
					}

				}
			}
			buf.WriteString(`]`)
		} else {
			buf.WriteString(`null`)
		}
		buf.WriteByte(',')
	}
	buf.Rewind(1)
	buf.WriteByte('}')
	return nil
}

const (
	ffj_t_ClusterStatus_DrainingMachinebase = iota
	ffj_t_ClusterStatus_DrainingMachineno_such_key

	ffj_t_ClusterStatus_DrainingMachine_Id

	ffj_t_ClusterStatus_DrainingMachine_Statuses
)

var ffj_key_ClusterStatus_DrainingMachine_Id = []byte("id")

var ffj_key_ClusterStatus_DrainingMachine_Statuses = []byte("statuses")

func (uj *ClusterStatus_DrainingMachine) UnmarshalJSON(input []byte) error {
	fs := fflib.NewFFLexer(input)
	return uj.UnmarshalJSONFFLexer(fs, fflib.FFParse_map_start)
}

func (uj *ClusterStatus_DrainingMachine) UnmarshalJSONFFLexer(fs *fflib.FFLexer, state fflib.FFParseState) error {
	var err error = nil
	currentKey := ffj_t_ClusterStatus_DrainingMachinebase
	_ = currentKey
	tok := fflib.FFTok_init
	wantedTok := fflib.FFTok_init

mainparse:
	for {
		tok = fs.Scan()
		//	println(fmt.Sprintf("debug: tok: %v  state: %v", tok, state))
		if tok == fflib.FFTok_error {
			goto tokerror
		}

		switch state {

		case fflib.FFParse_map_start:
			if tok != fflib.FFTok_left_bracket {
				wantedTok = fflib.FFTok_left_bracket
				goto wrongtokenerror
			}
			state = fflib.FFParse_want_key
			continue

		case fflib.FFParse_after_value:
			if tok == fflib.FFTok_comma {
				state = fflib.FFParse_want_key
			} else if tok == fflib.FFTok_right_bracket {
				goto done
			} else {
				wantedTok = fflib.FFTok_comma
				goto wrongtokenerror
			}

		case fflib.FFParse_want_key:
			// json {} ended. goto exit. woo.
			if tok == fflib.FFTok_right_bracket {
				goto done
			}
			if tok != fflib.FFTok_string {
				wantedTok = fflib.FFTok_string
				goto wrongtokenerror
			}

			kn := fs.Output.Bytes()
			if len(kn) <= 0 {
				// "" case. hrm.
				currentKey = ffj_t_ClusterStatus_DrainingMachineno_such_key
				state = fflib.FFParse_want_colon
				goto mainparse
			} else {
				switch kn[0] {

				case 'i':

					if bytes.Equal(ffj_key_ClusterStatus_DrainingMachine_Id, kn) {
						currentKey = ffj_t_ClusterStatus_DrainingMachine_Id
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 's':

					if bytes.Equal(ffj_key_ClusterStatus_DrainingMachine_Statuses, kn) {
						currentKey = ffj_t_ClusterStatus_DrainingMachine_Statuses
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				}

				if fflib.EqualFoldRight(ffj_key_ClusterStatus_DrainingMachine_Statuses, kn) {
					currentKey = ffj_t_ClusterStatus_DrainingMachine_Statuses
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_ClusterStatus_DrainingMachine_Id, kn) {
					currentKey = ffj_t_ClusterStatus_DrainingMachine_Id
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				currentKey = ffj_t_ClusterStatus_DrainingMachineno_such_key
				state = fflib.FFParse_want_colon
				goto mainparse
			}

		case fflib.FFParse_want_colon:
			if tok != fflib.FFTok_colon {
				wantedTok = fflib.FFTok_colon
				goto wrongtokenerror
			}
			state = fflib.FFParse_want_value
			continue
		case fflib.FFParse_want_value:

			if tok == fflib.FFTok_left_brace || tok == fflib.FFTok_left_bracket || tok == fflib.FFTok_integer || tok == fflib.FFTok_double || tok == fflib.FFTok_string || tok == fflib.FFTok_bool || tok == fflib.FFTok_null {
				switch currentKey {

				case ffj_t_ClusterStatus_DrainingMachine_Id:
					goto handle_Id

				case ffj_t_ClusterStatus_DrainingMachine_Statuses:
					goto handle_Statuses

				case ffj_t_ClusterStatus_DrainingMachineno_such_key:
					err = fs.SkipField(tok)
					if err != nil {
						return fs.WrapErr(err)
					}
					state = fflib.FFParse_after_value
					goto mainparse
				}
			} else {
				goto wantedvalue
			}
		}
	}

handle_Id:

	/* handler: uj.Id type=mesos.MachineID kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

			uj.Id = nil

			state = fflib.FFParse_after_value
			goto mainparse
		}

		if uj.Id == nil {
			uj.Id = new(mesos.MachineID)
		}

		err = uj.Id.UnmarshalJSONFFLexer(fs, fflib.FFParse_want_key)
		if err != nil {
			return err
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Statuses:

	/* handler: uj.Statuses type=[]*allocator.InverseOfferStatus kind=slice quoted=false*/

	{

		{
			if tok != fflib.FFTok_left_brace && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for ", tok))
			}
		}

		if tok == fflib.FFTok_null {
			uj.Statuses = nil
		} else {

			uj.Statuses = []*allocator.InverseOfferStatus{}

			wantVal := true

			for {

				var tmp_uj__Statuses *allocator.InverseOfferStatus

				tok = fs.Scan()
				if tok == fflib.FFTok_error {
					goto tokerror
				}
				if tok == fflib.FFTok_right_brace {
					break
				}

				if tok == fflib.FFTok_comma {
					if wantVal == true {
						// TODO(pquerna): this isn't an ideal error message, this handles
						// things like [,,,] as an array value.
						return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
					}
					continue
				} else {
					wantVal = true
				}

				/* handler: tmp_uj__Statuses type=*allocator.InverseOfferStatus kind=ptr quoted=false*/

				{
					if tok == fflib.FFTok_null {

						tmp_uj__Statuses = nil

						state = fflib.FFParse_after_value
						goto mainparse
					}

					if tmp_uj__Statuses == nil {
						tmp_uj__Statuses = new(allocator.InverseOfferStatus)
					}

					err = tmp_uj__Statuses.UnmarshalJSONFFLexer(fs, fflib.FFParse_want_key)
					if err != nil {
						return err
					}
					state = fflib.FFParse_after_value
				}

				uj.Statuses = append(uj.Statuses, tmp_uj__Statuses)

				wantVal = false
			}
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

wantedvalue:
	return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
wrongtokenerror:
	return fs.WrapErr(fmt.Errorf("ffjson: wanted token: %v, but got token: %v output=%s", wantedTok, tok, fs.Output.String()))
tokerror:
	if fs.BigError != nil {
		return fs.WrapErr(fs.BigError)
	}
	err = fs.Error.ToError()
	if err != nil {
		return fs.WrapErr(err)
	}
	panic("ffjson-generated: unreachable, please report bug.")
done:
	return nil
}

func (mj *Schedule) MarshalJSON() ([]byte, error) {
	var buf fflib.Buffer
	if mj == nil {
		buf.WriteString("null")
		return buf.Bytes(), nil
	}
	err := mj.MarshalJSONBuf(&buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
func (mj *Schedule) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	if mj == nil {
		buf.WriteString("null")
		return nil
	}
	var err error
	var obj []byte
	_ = obj
	_ = err
	buf.WriteString(`{ `)
	if len(mj.Windows) != 0 {
		buf.WriteString(`"windows":`)
		if mj.Windows != nil {
			buf.WriteString(`[`)
			for i, v := range mj.Windows {
				if i != 0 {
					buf.WriteString(`,`)
				}

				{

					if v == nil {
						buf.WriteString("null")
						return nil
					}

					err = v.MarshalJSONBuf(buf)
					if err != nil {
						return err
					}

				}
			}
			buf.WriteString(`]`)
		} else {
			buf.WriteString(`null`)
		}
		buf.WriteByte(',')
	}
	buf.Rewind(1)
	buf.WriteByte('}')
	return nil
}

const (
	ffj_t_Schedulebase = iota
	ffj_t_Scheduleno_such_key

	ffj_t_Schedule_Windows
)

var ffj_key_Schedule_Windows = []byte("windows")

func (uj *Schedule) UnmarshalJSON(input []byte) error {
	fs := fflib.NewFFLexer(input)
	return uj.UnmarshalJSONFFLexer(fs, fflib.FFParse_map_start)
}

func (uj *Schedule) UnmarshalJSONFFLexer(fs *fflib.FFLexer, state fflib.FFParseState) error {
	var err error = nil
	currentKey := ffj_t_Schedulebase
	_ = currentKey
	tok := fflib.FFTok_init
	wantedTok := fflib.FFTok_init

mainparse:
	for {
		tok = fs.Scan()
		//	println(fmt.Sprintf("debug: tok: %v  state: %v", tok, state))
		if tok == fflib.FFTok_error {
			goto tokerror
		}

		switch state {

		case fflib.FFParse_map_start:
			if tok != fflib.FFTok_left_bracket {
				wantedTok = fflib.FFTok_left_bracket
				goto wrongtokenerror
			}
			state = fflib.FFParse_want_key
			continue

		case fflib.FFParse_after_value:
			if tok == fflib.FFTok_comma {
				state = fflib.FFParse_want_key
			} else if tok == fflib.FFTok_right_bracket {
				goto done
			} else {
				wantedTok = fflib.FFTok_comma
				goto wrongtokenerror
			}

		case fflib.FFParse_want_key:
			// json {} ended. goto exit. woo.
			if tok == fflib.FFTok_right_bracket {
				goto done
			}
			if tok != fflib.FFTok_string {
				wantedTok = fflib.FFTok_string
				goto wrongtokenerror
			}

			kn := fs.Output.Bytes()
			if len(kn) <= 0 {
				// "" case. hrm.
				currentKey = ffj_t_Scheduleno_such_key
				state = fflib.FFParse_want_colon
				goto mainparse
			} else {
				switch kn[0] {

				case 'w':

					if bytes.Equal(ffj_key_Schedule_Windows, kn) {
						currentKey = ffj_t_Schedule_Windows
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				}

				if fflib.EqualFoldRight(ffj_key_Schedule_Windows, kn) {
					currentKey = ffj_t_Schedule_Windows
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				currentKey = ffj_t_Scheduleno_such_key
				state = fflib.FFParse_want_colon
				goto mainparse
			}

		case fflib.FFParse_want_colon:
			if tok != fflib.FFTok_colon {
				wantedTok = fflib.FFTok_colon
				goto wrongtokenerror
			}
			state = fflib.FFParse_want_value
			continue
		case fflib.FFParse_want_value:

			if tok == fflib.FFTok_left_brace || tok == fflib.FFTok_left_bracket || tok == fflib.FFTok_integer || tok == fflib.FFTok_double || tok == fflib.FFTok_string || tok == fflib.FFTok_bool || tok == fflib.FFTok_null {
				switch currentKey {

				case ffj_t_Schedule_Windows:
					goto handle_Windows

				case ffj_t_Scheduleno_such_key:
					err = fs.SkipField(tok)
					if err != nil {
						return fs.WrapErr(err)
					}
					state = fflib.FFParse_after_value
					goto mainparse
				}
			} else {
				goto wantedvalue
			}
		}
	}

handle_Windows:

	/* handler: uj.Windows type=[]*maintenance.Window kind=slice quoted=false*/

	{

		{
			if tok != fflib.FFTok_left_brace && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for ", tok))
			}
		}

		if tok == fflib.FFTok_null {
			uj.Windows = nil
		} else {

			uj.Windows = []*Window{}

			wantVal := true

			for {

				var tmp_uj__Windows *Window

				tok = fs.Scan()
				if tok == fflib.FFTok_error {
					goto tokerror
				}
				if tok == fflib.FFTok_right_brace {
					break
				}

				if tok == fflib.FFTok_comma {
					if wantVal == true {
						// TODO(pquerna): this isn't an ideal error message, this handles
						// things like [,,,] as an array value.
						return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
					}
					continue
				} else {
					wantVal = true
				}

				/* handler: tmp_uj__Windows type=*maintenance.Window kind=ptr quoted=false*/

				{
					if tok == fflib.FFTok_null {

						tmp_uj__Windows = nil

						state = fflib.FFParse_after_value
						goto mainparse
					}

					if tmp_uj__Windows == nil {
						tmp_uj__Windows = new(Window)
					}

					err = tmp_uj__Windows.UnmarshalJSONFFLexer(fs, fflib.FFParse_want_key)
					if err != nil {
						return err
					}
					state = fflib.FFParse_after_value
				}

				uj.Windows = append(uj.Windows, tmp_uj__Windows)

				wantVal = false
			}
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

wantedvalue:
	return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
wrongtokenerror:
	return fs.WrapErr(fmt.Errorf("ffjson: wanted token: %v, but got token: %v output=%s", wantedTok, tok, fs.Output.String()))
tokerror:
	if fs.BigError != nil {
		return fs.WrapErr(fs.BigError)
	}
	err = fs.Error.ToError()
	if err != nil {
		return fs.WrapErr(err)
	}
	panic("ffjson-generated: unreachable, please report bug.")
done:
	return nil
}

func (mj *Window) MarshalJSON() ([]byte, error) {
	var buf fflib.Buffer
	if mj == nil {
		buf.WriteString("null")
		return buf.Bytes(), nil
	}
	err := mj.MarshalJSONBuf(&buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
func (mj *Window) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	if mj == nil {
		buf.WriteString("null")
		return nil
	}
	var err error
	var obj []byte
	_ = obj
	_ = err
	buf.WriteString(`{ `)
	if len(mj.MachineIds) != 0 {
		buf.WriteString(`"machine_ids":`)
		if mj.MachineIds != nil {
			buf.WriteString(`[`)
			for i, v := range mj.MachineIds {
				if i != 0 {
					buf.WriteString(`,`)
				}

				{

					if v == nil {
						buf.WriteString("null")
						return nil
					}

					err = v.MarshalJSONBuf(buf)
					if err != nil {
						return err
					}

				}
			}
			buf.WriteString(`]`)
		} else {
			buf.WriteString(`null`)
		}
		buf.WriteByte(',')
	}
	if mj.Unavailability != nil {
		if true {
			buf.WriteString(`"unavailability":`)

			{

				err = mj.Unavailability.MarshalJSONBuf(buf)
				if err != nil {
					return err
				}

			}
			buf.WriteByte(',')
		}
	}
	buf.Rewind(1)
	buf.WriteByte('}')
	return nil
}

const (
	ffj_t_Windowbase = iota
	ffj_t_Windowno_such_key

	ffj_t_Window_MachineIds

	ffj_t_Window_Unavailability
)

var ffj_key_Window_MachineIds = []byte("machine_ids")

var ffj_key_Window_Unavailability = []byte("unavailability")

func (uj *Window) UnmarshalJSON(input []byte) error {
	fs := fflib.NewFFLexer(input)
	return uj.UnmarshalJSONFFLexer(fs, fflib.FFParse_map_start)
}

func (uj *Window) UnmarshalJSONFFLexer(fs *fflib.FFLexer, state fflib.FFParseState) error {
	var err error = nil
	currentKey := ffj_t_Windowbase
	_ = currentKey
	tok := fflib.FFTok_init
	wantedTok := fflib.FFTok_init

mainparse:
	for {
		tok = fs.Scan()
		//	println(fmt.Sprintf("debug: tok: %v  state: %v", tok, state))
		if tok == fflib.FFTok_error {
			goto tokerror
		}

		switch state {

		case fflib.FFParse_map_start:
			if tok != fflib.FFTok_left_bracket {
				wantedTok = fflib.FFTok_left_bracket
				goto wrongtokenerror
			}
			state = fflib.FFParse_want_key
			continue

		case fflib.FFParse_after_value:
			if tok == fflib.FFTok_comma {
				state = fflib.FFParse_want_key
			} else if tok == fflib.FFTok_right_bracket {
				goto done
			} else {
				wantedTok = fflib.FFTok_comma
				goto wrongtokenerror
			}

		case fflib.FFParse_want_key:
			// json {} ended. goto exit. woo.
			if tok == fflib.FFTok_right_bracket {
				goto done
			}
			if tok != fflib.FFTok_string {
				wantedTok = fflib.FFTok_string
				goto wrongtokenerror
			}

			kn := fs.Output.Bytes()
			if len(kn) <= 0 {
				// "" case. hrm.
				currentKey = ffj_t_Windowno_such_key
				state = fflib.FFParse_want_colon
				goto mainparse
			} else {
				switch kn[0] {

				case 'm':

					if bytes.Equal(ffj_key_Window_MachineIds, kn) {
						currentKey = ffj_t_Window_MachineIds
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'u':

					if bytes.Equal(ffj_key_Window_Unavailability, kn) {
						currentKey = ffj_t_Window_Unavailability
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				}

				if fflib.SimpleLetterEqualFold(ffj_key_Window_Unavailability, kn) {
					currentKey = ffj_t_Window_Unavailability
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffj_key_Window_MachineIds, kn) {
					currentKey = ffj_t_Window_MachineIds
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				currentKey = ffj_t_Windowno_such_key
				state = fflib.FFParse_want_colon
				goto mainparse
			}

		case fflib.FFParse_want_colon:
			if tok != fflib.FFTok_colon {
				wantedTok = fflib.FFTok_colon
				goto wrongtokenerror
			}
			state = fflib.FFParse_want_value
			continue
		case fflib.FFParse_want_value:

			if tok == fflib.FFTok_left_brace || tok == fflib.FFTok_left_bracket || tok == fflib.FFTok_integer || tok == fflib.FFTok_double || tok == fflib.FFTok_string || tok == fflib.FFTok_bool || tok == fflib.FFTok_null {
				switch currentKey {

				case ffj_t_Window_MachineIds:
					goto handle_MachineIds

				case ffj_t_Window_Unavailability:
					goto handle_Unavailability

				case ffj_t_Windowno_such_key:
					err = fs.SkipField(tok)
					if err != nil {
						return fs.WrapErr(err)
					}
					state = fflib.FFParse_after_value
					goto mainparse
				}
			} else {
				goto wantedvalue
			}
		}
	}

handle_MachineIds:

	/* handler: uj.MachineIds type=[]*mesos.MachineID kind=slice quoted=false*/

	{

		{
			if tok != fflib.FFTok_left_brace && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for ", tok))
			}
		}

		if tok == fflib.FFTok_null {
			uj.MachineIds = nil
		} else {

			uj.MachineIds = []*mesos.MachineID{}

			wantVal := true

			for {

				var tmp_uj__MachineIds *mesos.MachineID

				tok = fs.Scan()
				if tok == fflib.FFTok_error {
					goto tokerror
				}
				if tok == fflib.FFTok_right_brace {
					break
				}

				if tok == fflib.FFTok_comma {
					if wantVal == true {
						// TODO(pquerna): this isn't an ideal error message, this handles
						// things like [,,,] as an array value.
						return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
					}
					continue
				} else {
					wantVal = true
				}

				/* handler: tmp_uj__MachineIds type=*mesos.MachineID kind=ptr quoted=false*/

				{
					if tok == fflib.FFTok_null {

						tmp_uj__MachineIds = nil

						state = fflib.FFParse_after_value
						goto mainparse
					}

					if tmp_uj__MachineIds == nil {
						tmp_uj__MachineIds = new(mesos.MachineID)
					}

					err = tmp_uj__MachineIds.UnmarshalJSONFFLexer(fs, fflib.FFParse_want_key)
					if err != nil {
						return err
					}
					state = fflib.FFParse_after_value
				}

				uj.MachineIds = append(uj.MachineIds, tmp_uj__MachineIds)

				wantVal = false
			}
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Unavailability:

	/* handler: uj.Unavailability type=mesos.Unavailability kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

			uj.Unavailability = nil

			state = fflib.FFParse_after_value
			goto mainparse
		}

		if uj.Unavailability == nil {
			uj.Unavailability = new(mesos.Unavailability)
		}

		err = uj.Unavailability.UnmarshalJSONFFLexer(fs, fflib.FFParse_want_key)
		if err != nil {
			return err
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

wantedvalue:
	return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
wrongtokenerror:
	return fs.WrapErr(fmt.Errorf("ffjson: wanted token: %v, but got token: %v output=%s", wantedTok, tok, fs.Output.String()))
tokerror:
	if fs.BigError != nil {
		return fs.WrapErr(fs.BigError)
	}
	err = fs.Error.ToError()
	if err != nil {
		return fs.WrapErr(err)
	}
	panic("ffjson-generated: unreachable, please report bug.")
done:
	return nil
}
