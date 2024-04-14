package service

func Scan2Epub(chaps []string) error {
	for _, chap := range chaps {
		if err := downloadChap(chap); err != nil {
			return err
		}
		// if err := convertChap(chap); err != nil {
		// 	return err
		// }
	}
	return nil
}
