package validation

import (
	"fmt"
	"regexp"
)

type Rule func(key string, value interface{}) error

type Rules []Rule

type Validator struct {
	rules Rules
}

func (v *Validator) Add(rule Rule) {
	v.rules = append(v.rules, rule)
}

func (v *Validator) Validate(data map[string]interface{}) []error {
	var errors []error
	for _, rule := range v.rules {
		for key, value := range data {
			if err := rule(key, value); err != nil {
				errors = append(errors, err)
			}
		}
	}
	return errors
}

func ValidateNIC(key string, value interface{}) error {
	str, ok := value.(string)
	if !ok {
		return fmt.Errorf("%s is not string", key)
	}
	nicRegex := `^(([5-9]{1}[0-9]{1}[0-3,5-8]{1}[0-9]{6}[vVxX])|([1-2]{1}[0-9]{2}[0-3,5-8]{1}[0-9]{7}))$`
	matched, err := regexp.MatchString(nicRegex, str)
	if err != nil {
		return fmt.Errorf("error occurred while validating %s: %v", key, err)
	}
	if !matched {
		return fmt.Errorf("%s is not a nic", key)
	}
	return nil
}

func ValidateEmail(key string, value interface{}) error {
	str, ok := value.(string)
	if !ok {
		return fmt.Errorf("%s is not a string", key)
	}
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	match, err := regexp.MatchString(emailRegex, str)
	if err != nil {
		return fmt.Errorf("error occurred while validating %s: %v", key, err)
	}
	if !match {
		return fmt.Errorf("%s is not a valid email address", key)
	}
	return nil
}
