package models

import "testing"

// TestModelosSeRegistran comprueba que el ORM acepta las definiciones de este
// paquete. Basta con que el test arranque: el init() de cada archivo llama a
// orm.RegisterModel, que entra en panico si un tipo de campo no esta
// soportado.
//
// Importa en concreto por los campos *int de Pagos (id_contrataciones,
// id_tipo_cuenta): son columnas anulables que no son relaciones, y hacia falta
// confirmar que Beego los admite como puntero en vez de int.
func TestModelosSeRegistran(t *testing.T) {
	p := &Pagos{}
	if p.TableName() != "pagos" {
		t.Fatalf("TableName = %q", p.TableName())
	}

	// Un puntero nulo representa "sin contratacion" y "sin tipo de cuenta".
	if p.IdContrataciones != nil {
		t.Error("IdContrataciones deberia arrancar nulo")
	}
	if p.IdTipoCuenta != nil {
		t.Error("IdTipoCuenta deberia arrancar nulo")
	}

	// Y con valor, se puede distinguir del cero.
	cero := 0
	p.IdContrataciones = &cero
	if p.IdContrataciones == nil || *p.IdContrataciones != 0 {
		t.Error("un puntero a cero deberia distinguirse de nulo")
	}
}
