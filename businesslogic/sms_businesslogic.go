package businesslogic

import "fmt"

func CreateSmsWording(nameRef string, name string) string {
	return fmt.Sprintf("ขอบคุณคุณ%s ที่แนะนำคูณ%s ให้มาซื้อประกันกับเรา สอบถาม1557", nameRef, name)
}

func CreateSmsWordingWithoutCustomer(name string) string {
	return fmt.Sprintf("ขอบคุณที่แนะนำคูณ%s ให้มาซื้อประกันกับเรา สอบถาม1557", name)
}
