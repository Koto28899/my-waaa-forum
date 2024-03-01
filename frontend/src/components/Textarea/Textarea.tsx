type TextareaPropsRequired = {
  id: string,
  name: string,
  placeholder: string,
  minLength: number,
  maxLength: number,
  required: boolean,
  rows: number,
  cols: number,
}

function Textarea({ propsRequired }: { propsRequired: TextareaPropsRequired}) {
  return (
    <div className="textarea-group">
      <textarea className="textarea" 
        id={propsRequired.id}
        placeholder={propsRequired.placeholder}
        name={propsRequired.id} 
        rows={propsRequired.rows} 
        cols={propsRequired.cols}
        minLength={propsRequired.minLength}
        maxLength={propsRequired.maxLength}
        required={propsRequired.required}/>
      <label htmlFor={propsRequired.id} className="textarea-label">
        {propsRequired.placeholder}
      </label>
    </div>
  )
}

export type {
  TextareaPropsRequired
}

export default Textarea;