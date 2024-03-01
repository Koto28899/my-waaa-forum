'use client'

import React, { useState } from "react";
import Input, { 
  ColorValidInput, 
  ColorInvalidInput, 
} from ".";

function ManagerIDInput() {

  const [inputPlaceHolder, setInputPlaceHolder] = useState("管理员 ID")
  var timeoutID: NodeJS.Timeout;

  function IDOnChangeHandler(changeEvent: React.ChangeEvent<Element>) {
    setInputPlaceHolder("读取中...");
    const element = changeEvent.target as HTMLInputElement;
    element.style.borderBottomColor = ColorInvalidInput;
    element.oninput = function () {
      if (element.value.length > 5) {
        element.value = element.value.slice(0,5);
        return
      }
      if (element.valueAsNumber < 1) {
        element.valueAsNumber = 1;
        setInputPlaceHolder("管理员 ID 最小为 1")
        return
      }
    }
    clearTimeout(timeoutID);
    if (element.value == "") {  
      return 
    }
    timeoutID = setTimeout(function() {
      fetch(`http://localhost:8080/api/role/basis?id=${(element).value}`, {
        method: "GET"
      }).then(response => {
        if (!response.ok || response.status != 200) {
          return Promise.reject({
            code: response.status
          })
        }
        response.json().then(json => {
          if (json.type == "user") {
            return Promise.reject({
              code: 401
            })
          }
        })
        element.style.borderBottomColor = ColorValidInput
        setInputPlaceHolder("ID 有效")
      }).catch(err => {
        switch (err.code) {
          case 400: {
            setInputPlaceHolder("无效请求体");
            break;
          }
          case 401: {
            setInputPlaceHolder("提供的用户 ID 无效，因为他不是管理员");
            break;
          }
          case 404: {
            setInputPlaceHolder("提供的用户 ID 无效，用户不存在");
            break;
          }
          default: {
            setInputPlaceHolder("未知错误");
          }
        }
      })
    }, 1000)
  }

  return (
    <Input 
      id="__ManagerID"
      type="number"
      name="__ManagerID"
      placeholder={inputPlaceHolder}
      minLength={1}
      maxLength={999999}
      required={true}
      onChange={IDOnChangeHandler}
    />
  )
}

export default ManagerIDInput;