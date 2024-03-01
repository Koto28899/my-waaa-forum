'use client'

import { useState } from "react";
import Input, { 
  ColorValidInput, 
  ColorInvalidInput, 
} from ".";

function NewSectionNameInput() {

  const [_placeholder, _setPlaceholder] = useState("全新的论坛名")
  var timeoutID: NodeJS.Timeout;

  function SectionNameOnChangeHandler(changeEvent: React.ChangeEvent<Element>) {
    _setPlaceholder("读取中...")
    const element = changeEvent.target as HTMLInputElement
    element.style.borderBottomColor = ColorInvalidInput
    clearTimeout(timeoutID)
    if (element.value == "") { 
      return 
    }
    timeoutID = setTimeout(function() {
      fetch(`http://localhost:8080/api/section/only?sectionName=${(element).value}`, {
        method: "GET"
      }).then(response => {
        if (!response.ok) {
          return Promise.reject({
            code: response.status
          })
        } else {
          if (response.status == 200) {
            _setPlaceholder("版块已存在");
            return;
          }
        }
      }).catch(err => {
        switch(err.code) {
          case 400: {
            _setPlaceholder("请求体错误")
            break
          }
          case 404: {
            element.style.borderBottomColor = ColorValidInput
            _setPlaceholder("版块名可使用")
            break
          }
          default: {
            _setPlaceholder("未知错误")
          }
        }
      })
    }, 1000)
  }

  return (
    <Input
      id="__name"
      type="text"
      name="__name"
      placeholder={_placeholder}
      minLength={2}
      maxLength={8}
      required={true}
      onChange={SectionNameOnChangeHandler}
    />
  )
}

export default NewSectionNameInput;