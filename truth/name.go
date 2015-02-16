package truth

func (t *truthControl) GatherName() (TruthCollection, []Truth, error) {
	return truthHook("name"), nil, nil
}

func (t *truthControl) GatherAge() (TruthCollection, []Truth, error) {
	return truthHook("age"), nil, nil
}

func (t *truthControl) GatherSalary() (TruthCollection, []Truth, error) {
	return truthHook("salary"), nil, nil
}

func (t *truthControl) GatherSpouse() (TruthCollection, []Truth, error) {
	return truthHook("spouse"), nil, nil
}

func (t *truthControl) GatherKids() (TruthCollection, []Truth, error) {
	return truthHook("kids"), nil, nil
}
