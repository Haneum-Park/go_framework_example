# 명명규칙

## MixedCaps or mixedCaps

- 첫 문자가 대문자 혹은 소문자에 따라 접근 지정자가 달라짐
- Public일 경우 PascalCase를, Private일 경우 camelCase를 사용

## Package names

- 패키지 이름은 반드시 lowercase 작성되며 underscore 또는 mixedCaps를 금함
- 패키지 이름은 반드시 유일해야하며, "lower/case"로 import 하되, "case"로 접근한다.

## Interface names

- Interface 이름은,
  하나의 메소드만 있는 경우 MethodName + er로 하며,
  복수개의 메소드가 있는 경우 MethodName + ers로 한다.

## Variable names

- camelCase의 mixedCaps로 사용하며, 최대한 짧게, 길다면 약어로 사용한다.

## Unique names

- HTTP, API와 같은 acronym의 경우에는 그냥 uppercase를 유지한다.
