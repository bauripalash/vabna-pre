![Vabna Programming Language](./images/vabna_cover.jpg)

## Introduction
Vabna is a interpreted dynamically typed programming language for programming in Bengali language. Vabna is designed with Bengali with mind but can be also used with English, in fact the implementation is so easy to modify that it can be used to program in any languages with very little change to the source code of vabna interpreter.

## Language Features

###  Data Types:
* Strings : `"পলাশ বাউরি"` , `"ভাবনা"`...
* Integers: `99999` , `1234567890`
* Dictionaries/Hashmap : `{ "নাম": "পলাশ", "বয়স" : 20  }`
* Arrays: `["রবিবার", "সোমবার" , 21 , 22 , 23]`   
* Booleans: `সত্য`, `মিথ্যা`

### Functions:
* Example: 
```go
ধরি ঘুমানো = একটি কাজ(নায়ক){
    দেখাও(নায়ক + " ঘুমোচ্ছে!");
}; 

ঘুমানো("পলাশ বাউরি")
```
```
Output: পলাশ বাউরি ঘুমোচ্ছে!
```
### Assignments:
* Examples:
```go
ধরি মাস = "বৈশাখ";
ধরি আজ_কি_ছুটি = মিথ্যা; 
```

## Project Status:
> **Alpha** (*Under Heavy Development*) 

## License:
> GNU GPL

## Special Thanks,
* Thorsten Ball for writing this amazing book "
Writing An Interpreter In Go
"