package main


func main() {
    app := &AppSpec{
        Name:     "guestbook",
        Replicas: 2,
		Containers: []string{
            "nginx",
            "redis",
			"mongodb",
        },
        Labels: map[string]string{
            "app": "guestbook",
            "env": "dev",
        },
    }

    app.reconcile()
}
	
