import React from 'react';
import { useState } from 'react';
import { Link } from 'react-router-dom';
import "./styles/Auth.scss";
import { register } from './service/service-auth';
import Input from './components/Input';
import "./styles/Component.scss";
import { useNavigate } from 'react-router-dom';
const Register = () =>
{
    const [email, setEmail ] = useState( "" );
    const [loginName, setLoginName] = useState("");
    const [password, setPassword] = useState("");
    const [ confirm, setConfirm ] = useState( "" );

    const [emailInvalid, setEmailInvalid ] = useState( false );
    const [loginNameInvalid, setLoginNameInvalid] = useState(false);
    const [passwordInvalid, setPasswordInvalid] = useState(false);
    const [ confirmInvalid, setConfirmInvalid ] = useState( false );
    
    const navigate = useNavigate()

    const handleSubmit = e =>
    {
        e.preventDefault();
        let invalid = false;
        const noEmptyString = "This field cannot be empty.";

        if ( password != confirm )
        {
            setConfirmInvalid( "Confirm password is not match." )
            invalid = true;
        }

        const checkEmail = (email) =>
        {
            const pattern = /^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$/;
            return pattern.test(email)
        }

        if ( !checkEmail( email ) )
        {
            setEmailInvalid( "Email is invalid." )
            invalid = true;
        }
        if(password.length <6)
        {
            setPasswordInvalid("Password must have more than 6 character.")
            invalid = true;
        }

        if ( email == "" )
        {
            setEmailInvalid( noEmptyString );
            invalid = true;
        }
        if( loginName =="" )
        {
            setLoginNameInvalid( noEmptyString );
            invalid = true;
        }
        if(password == "")
        {
            setPasswordInvalid( noEmptyString );
            invalid = true;
        }
        if(confirm == "")
        {
            setConfirmInvalid( noEmptyString );
            invalid = true;
        }
        

        

        if(invalid)
        {
            return;
        }

        
        register( { email: email, id: loginName, password: password } ).then( (result) =>
        {
            console.log( "sucess" );
            navigate( "/login" );
        } ).catch( (ex) =>
        {
            const message = ex.response.data;
            if ( message.emailExist )
            {
                setEmailInvalid( "Email exist" );
            }
            if ( message.nameExist )
            {
                setLoginNameInvalid( "Login name exist" );
            }
            
        })

    }

    return ( 
        <div className="auth">
            <div className="auth-container tag">
                <h2>Register</h2>
                <form onSubmit={handleSubmit}>
                    <Input fieldName="Email:" type="text" id="email" placeholder="example@gmail.com" value={email} onChange={e=>{setEmail(e.target.value)}} invalidText={emailInvalid} removeInvalidStatus={e=>setEmailInvalid("")} />
                    <Input fieldName="Login name:" type="text" id="login-name" placeholder="example" value={loginName} onChange={e=>setLoginName(e.target.value)} invalidText={loginNameInvalid} removeInvalidStatus={e=>setLoginNameInvalid("")} />
                    <Input fieldName="Mật khẩu:" type="password" id="password" placeholder="• • • • • • • •" value={password} onChange={e=>{setPassword(e.target.value)}} invalidText={passwordInvalid} removeInvalidStatus={e=>setPasswordInvalid("")} />
                    <Input fieldName="Nhập lại mật khẩu:" type="password" id="confirm" placeholder="• • • • • • • •" value={confirm} onChange={e=>{setConfirm(e.target.value)}} invalidText={confirmInvalid} removeInvalidStatus={e=>setConfirmInvalid("")} />
                    
                <span>Bạn đã có tài khoản? <Link to="/login" style={{textDecoration:"underline", fontWeight:"500"}}>đăng nhập</Link></span>
                <button type="submit">Đăng ký</button>
                </form>
            </div>
        </div>
     );
}
 
export default Register;