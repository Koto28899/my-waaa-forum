import React, { InputHTMLAttributes } from "react";

const Input = React.forwardRef(
  function(
    propsRequired: InputHTMLAttributes<HTMLInputElement>, 
    ref: React.ForwardedRef<HTMLInputElement>
    ) {
    return (
      <div className="input-group" key={propsRequired.id}>
        <input className="input" 
          accept={propsRequired.accept}
          alt={propsRequired.alt}
          autoComplete={propsRequired.autoComplete}
          capture={propsRequired.capture}
          checked={propsRequired.checked}
          disabled={propsRequired.disabled}
          enterKeyHint={propsRequired.enterKeyHint}
          form={propsRequired.form}
          formAction={propsRequired.formAction}
          formEncType={propsRequired.formEncType}
          formMethod={propsRequired.formMethod}
          formNoValidate={propsRequired.formNoValidate}
          formTarget={propsRequired.formTarget}
          height={propsRequired.height}
          list={propsRequired.list}
          max={propsRequired.max}
          maxLength={propsRequired.maxLength}
          min={propsRequired.min}
          minLength={propsRequired.minLength}
          multiple={propsRequired.multiple}
          name={propsRequired.name}
          pattern={propsRequired.pattern}
          placeholder={propsRequired.placeholder}
          readOnly={propsRequired.readOnly}
          required={propsRequired.required}
          size={propsRequired.size}
          src={propsRequired.src}
          step={propsRequired.step}
          type={propsRequired.type}
          value={propsRequired.value}
          width={propsRequired.width}
          onChange={propsRequired.onChange}

          ref={ref}
        />
        <label className="input-label" htmlFor={propsRequired.id}>
          {propsRequired.placeholder}
        </label>
      </div>
    )
  }
)

export default Input;