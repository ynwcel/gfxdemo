package svcx

func RunModeIsDebug() bool {
	if v := svcx_maps.Get(svcx_runmode_mapkey); v != nil {
		if v_bool, ok := v.(bool); ok && v_bool {
			return true
		}
	}
	return false
}
