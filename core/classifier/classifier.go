package classifier


type Classifier struct {
	Info map[string]string
}

func (cls *Classifier) State() string {
	stat, ok := cls.Info["state"]
	if ok {
		return stat
	}
	return "empty"
}


