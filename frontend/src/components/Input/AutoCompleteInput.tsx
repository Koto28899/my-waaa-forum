import React from "react";
import Input from "./Input";

interface AutoCompleteInputHTMLAttributes<HTMLInputElement> extends React.InputHTMLAttributes<HTMLInputElement> {
  suggestions: any[],
}

function AutoCompleteInput() {
  return (
    <Input 
      
    />
  )
}

export default AutoCompleteInput;