package pattern

import "fmt"

/*
	Реализовать паттерн «цепочка вызовов».
	Объяснить применимость паттерна, его плюсы и минусы,
	а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern
*/

func ExecuteCoPExample() {
	cashier := &Cashier{}

	pharmacy := &Pharmacy{}
	pharmacy.Refer(cashier)

	doctor := &Doctor{}
	doctor.Refer(pharmacy)

	reception := &Reception{}
	reception.Refer(doctor)

	patient := &Patient{
		name: "John Doe",
	}

	reception.Receive(patient)
}

type Polyclinic interface {
	Receive(*Patient)
	Refer(Polyclinic)
}

type Reception struct {
	next Polyclinic
}

func (r *Reception) Receive(p *Patient) {
	if p.registered {
		fmt.Println("Registration is already done")
		r.next.Receive(p)
		return
	}

	fmt.Println("Reception has completed a patient registration")
	r.next.Receive(p)
}

func (r *Reception) Refer(further Polyclinic) {
	r.next = further
}

type Doctor struct {
	next Polyclinic
}

func (d *Doctor) Receive(p *Patient) {
	if p.checkedByDoctor {
		fmt.Println("Doctor has already checked this patient")
		d.next.Receive(p)
		return
	}

	fmt.Println("Doctor has completed a parient checking")
	d.next.Receive(p)
}

func (d *Doctor) Refer(further Polyclinic) {
	d.next = further
}

type Pharmacy struct {
	next Polyclinic
}

func (ph *Pharmacy) Receive(p *Patient) {
	if p.providedWithMedicines {
		fmt.Println("Pharmacy has already prodived this patient with medicine")
		ph.next.Receive(p)
		return
	}

	fmt.Println("Pharmacy has completed a providing this patient with medicine")
	ph.next.Receive(p)
}

func (ph *Pharmacy) Refer(further Polyclinic) {
	ph.next = further
}

type Cashier struct {
	next Polyclinic
}

func (c *Cashier) Receive(p *Patient) {
	if p.paidTheBill {
		fmt.Println("Bill is already paid")
		return
	}

	fmt.Println("Cashier received the money from this patient")
}

func (c *Cashier) Refer(further Polyclinic) {
	c.next = further
}

type Patient struct {
	name                  string
	registered            bool
	checkedByDoctor       bool
	providedWithMedicines bool
	paidTheBill           bool
}
