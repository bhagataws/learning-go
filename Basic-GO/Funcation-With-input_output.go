package main

type AppSpec struct {
	Name     string
	Replicas int
}

func (a *AppSpec) GetName(name string, replicas int) (string, int, error) { //* menas Pointer Only, So that update the values of original struct,
	a.Name = name                  // Update by to pointer(Name Replicas will fillup from name , replicas  under the GetName(name string, replicas int)funcation )
	a.Replicas = replicas          // Update by to pointer(Name Replicas will fillup from name , replicas  under the GetName(name string, replicas int)funcation )
	return a.Name, a.Replicas, nil // It's return name and replicas in string format
}

func main() {
	app := &AppSpec{}               // We used Stucted Store it;s values inside us
	println(app.GetName("sonu", 3)) // Use the funcation GetName with required variables which is required to run GetName funcation
}
