package waveshare7in5v2

const (
	EPD_WIDTH  int = 800
	EPD_HEIGHT int = 480
)

const (
	PANEL_SETTING                byte = 0x00
	POWER_SETTING                byte = 0x01
	POWER_OFF                    byte = 0x02
	POWER_ON                     byte = 0x04
	BOOSTER_SOFT_START           byte = 0x06
	DEEP_SLEEP                   byte = 0x07
	DISPLAY_START_TRANSMISSION_1 byte = 0x10
	DISPLAY_REFRESH              byte = 0x12
	DISPLAY_START_TRANSMISSION_2 byte = 0x13
	DUAL_SPI                     byte = 0x15
	VCOM_DATA_INTERVAL_SETTING   byte = 0x50
	TCON                         byte = 0x60
	RESOLUTION_SETTING           byte = 0x61
	GET_STATUS                   byte = 0x71
	VCOM_DC                      byte = 0x82
)

const MAX_CHUNK_SIZE = 4096

const PIXEL_SIZE = 8

func getLutVcom() []byte {
	return []byte{
		0x0, 0xF, 0xF, 0x0, 0x0, 0x1,
		0x0, 0xF, 0x1, 0xF, 0x1, 0x2,
		0x0, 0xF, 0xF, 0x0, 0x0, 0x1,
		0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
		0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
		0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
		0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
	}
}

func getLutWw() []byte {
	return []byte{
		0x10, 0xF, 0xF, 0x0, 0x0, 0x1,
		0x84, 0xF, 0x1, 0xF, 0x1, 0x2,
		0x20, 0xF, 0xF, 0x0, 0x0, 0x1,
		0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
		0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
		0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
		0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
	}
}

func getLutBw() []byte {
	return []byte{
		0x10, 0xF, 0xF, 0x0, 0x0, 0x1,
		0x84, 0xF, 0x1, 0xF, 0x1, 0x2,
		0x20, 0xF, 0xF, 0x0, 0x0, 0x1,
		0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
		0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
		0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
		0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
	}
}

func getLutWb() []byte {
	return []byte{
		0x80, 0xF, 0xF, 0x0, 0x0, 0x1,
		0x84, 0xF, 0x1, 0xF, 0x1, 0x2,
		0x40, 0xF, 0xF, 0x0, 0x0, 0x1,
		0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
		0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
		0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
		0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
	}
}

func getLutBb() []byte {
	return []byte{
		0x80, 0xF, 0xF, 0x0, 0x0, 0x1,
		0x84, 0xF, 0x1, 0xF, 0x1, 0x2,
		0x40, 0xF, 0xF, 0x0, 0x0, 0x1,
		0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
		0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
		0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
		0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
	}
}

func getLutVcomFast() []byte {
	return []byte{
		//0x0, 0xF, 0xF, 0x0, 0x0, 0x1,
		//0x0, 0xF, 0x1, 0xF, 0x1, 0x2,
		//0x0, 0xF, 0xF, 0x0, 0x0, 0x1,
		//0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
		//0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
		//0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
		//0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
	}
}

func getLutWwFast() []byte {
	return []byte{
		//0x10, 0xF, 0xF, 0x0, 0x0, 0x1,
		//0x84, 0xF, 0x1, 0xF, 0x1, 0x2,
		//0x20, 0xF, 0xF, 0x0, 0x0, 0x1,
		//0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
		//0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
		//0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
		//0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
	}
}

func getLutBwFast() []byte {
	return []byte{
		//0x10, 0xF, 0xF, 0x0, 0x0, 0x1,
		//0x84, 0xF, 0x1, 0xF, 0x1, 0x2,
		//0x20, 0xF, 0xF, 0x0, 0x0, 0x1,
		//0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
		//0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
		//0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
		//0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
	}
}

func getLutWbFast() []byte {
	return []byte{
		//0x80, 0xF, 0xF, 0x0, 0x0, 0x1,
		//0x84, 0xF, 0x1, 0xF, 0x1, 0x2,
		//0x40, 0xF, 0xF, 0x0, 0x0, 0x1,
		//0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
		//0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
		//0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
		//0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
	}
}

func getLutBbFast() []byte {
	return []byte{
		//0x80, 0xF, 0xF, 0x0, 0x0, 0x1,
		//0x84, 0xF, 0x1, 0xF, 0x1, 0x2,
		//0x40, 0xF, 0xF, 0x0, 0x0, 0x1,
		//0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
		//0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
		//0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
		//0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
	}
}
