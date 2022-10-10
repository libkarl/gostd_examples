package interfaceex

import "fmt"


type Methodhandler interface {  
	GiveMeLenght() int
	CreateString()  string 	
}

type ObjT []string

func (o ObjT) GiveMeLenght() int {
	len := len(o)
	return len
}  

func (o ObjT) CreateString() string {
	var text string
	for i := 0; i < len(o); i++ {
		text += o[i]
	}
	return text
}

// advanced usecase 1
func ObjectTest1 () {
	pemp1 := Permanent{
		empId:    2,
		basicpay: 6000,
		pf:       30,
	}
	cemp1 := Contract{
		empId:    3,
		basicpay: 3000,
 	}
	pemp2 := Permanent{
		empId:    1,
		basicpay: 3000,
		pf:       30,
	}
	// if you want put just one item in interface -> in := SalaryCalculator(pemp1)
	in := []SalaryCalculator{pemp1, cemp1, pemp2}

	for _, val := range in {
		result, ok := val.(Permanent)
		fmt.Println(ok)
		if ok == true {
			salary := val.CalculateSalary()
			fmt.Println(result.empId, result.basicpay, result.pf, "calculated salary with pf: ", salary)
		} else {
			res, _ := val.(Contract)
			salary := res.CalculateSalary()
			fmt.Println(res.empId, res.basicpay, "calculated salary: ", salary)
		}
	}
	
}

// advanced use case 2
func ObjectTest2 () {
	pemp1 := Permanent{
		empId:    2,
		basicpay: 6000,
		pf:       30,
	}
	cemp1 := Contract{
		empId:    3,
		basicpay: 3000,
 	}
	pemp2 := Permanent{
		empId:    1,
		basicpay: 3000,
		pf:       30,
	}
	demp1 := Def{
		empID: 15,
	}
	in := []SalaryCalculator{pemp1, cemp1, pemp2, demp1}
	for _, val := range in {
		switch y := val.(type) {
		case Permanent:
			y.CalculateSalary()
			fmt.Println(y.empId, y.basicpay, y.pf)
		case Contract:
			y.CalculateSalary() 
			fmt.Println(y.empId, y.basicpay)
		default:
			fmt.Println(val)
		}
	}
	
	
	
}

// x := interfaceex.ObjT{"A", "B"}
func ObjectTest3(x ObjT) {
	fmt.Println(x.GiveMeLenght())
	fmt.Println(x.CreateString())
}

type SalaryCalculator interface {  
	CalculateSalary() int
    }
    
    type Permanent struct {  
	empId    int
	basicpay int
	pf       int
    }
    
    type Contract struct {  
	empId    int
	basicpay int
    }

    type Def struct {
	empID int
    }

    // def method
    func (d Def) CalculateSalary() int {
	return d.empID
    }
    
    //salary of permanent employee is the sum of basic pay and pf
    func (p Permanent) CalculateSalary() int {  
	return p.basicpay + p.pf
    }
    
    //salary of contract employee is the basic pay alone
    func (c Contract) CalculateSalary() int {  
	return c.basicpay
    }

