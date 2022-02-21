package count

import "testing"

func TestTambah(t *testing.T){
	result := Tambah(1,2);
	expected := 3;
	if result != expected {
		t.Errorf("Failed Result Not Match Expected");	
	}
}