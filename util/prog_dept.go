package util

// array of all double major program department IDs
/*
  28: AE	| 29: BSBE	| 30: CE	| 31: CHE	|
  32: CSE	| 33: EE	| 34: MSE	| 35: ME	|
  96: CHM	| 36: ECO	| 37: MTH	| 97: SDS	|
  98: PHY	|
*/
var doubleMajorProgramDepartmentIDs = []uint{28, 29, 30, 31, 32, 33, 34, 35, 96, 36, 37, 97, 98}

func IsDoubleMajor(programDepartmentID uint) bool {
	for _, id := range doubleMajorProgramDepartmentIDs {
		if id == programDepartmentID {
			return true
		}
	}
	return false
}
