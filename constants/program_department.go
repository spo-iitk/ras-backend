package constants

type ProgramID uint

type DepartmentID uint

type ProgramDepartment struct {
	ID   uint
	Prog ProgramID
	Dept DepartmentID
}

const (
	BT    ProgramID = 0
	BS    ProgramID = 1
	MTech ProgramID = 2
	MSc   ProgramID = 3
	MDes  ProgramID = 4
)

const (
	CSE  DepartmentID = 0
	EE   DepartmentID = 1
	MTH  DepartmentID = 2
	CHE  DepartmentID = 3
	ME   DepartmentID = 4
	CE   DepartmentID = 5
	MSE  DepartmentID = 6
	PHY  DepartmentID = 7
	CHM  DepartmentID = 8
	ECO  DepartmentID = 9
	BSBE DepartmentID = 10
	PSE  DepartmentID = 11
)

var (
	BS_MTH ProgramDepartment = ProgramDepartment{ID: 0, Prog: BS, Dept: MTH}
	BT_CSE ProgramDepartment = ProgramDepartment{ID: 1, Prog: BT, Dept: CSE}
)

func GetProgram(program ProgramID) string {
	switch program {
	case BS:
		return "BS"
	case BT:
		return "BT"
	default:
		return ""
	}
}

func GetDepartment(department DepartmentID) string {
	switch department {
	case CSE:
		return "CSE"
	case MTH:
		return "MTH"
	default:
		return ""
	}
}

func GetProgramDepartment(programDepartmentID uint) ProgramDepartment {
	switch programDepartmentID {
	case 0:
		return BS_MTH
	case 1:
		return BT_CSE
	default:
		return ProgramDepartment{}
	}
}
