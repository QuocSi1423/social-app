import React from "react";

const Input = (props) =>
{
    const { id, fieldName, placeholder, type, value, onChange, invalidText, removeInvalidStatus } = props;

    return (
        <div className={invalidText == ""? "input-invalid":"input-invalid invalid"}>
            <label htmlFor={id}>{ fieldName}</label>
            <input placeholder={ placeholder } id={ id } type={ type } value={ value } onChange={ onChange } onFocus={ removeInvalidStatus } />
            {
                !invalidText == "" ?
                    <div>{ invalidText }</div>
                    :
                    <div></div>
                    
            }
        </div>
    )
}

export default Input