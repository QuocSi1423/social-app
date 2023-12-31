import React from "react";  
import { Link, useNavigate } from "react-router-dom";
import { useState } from "react";
import Input from "./components/Input";
import { login } from "./service/service-auth";
import { useDispatch } from "react-redux";
import { updateUser } from "./redux/userSlice";
import "./styles/Auth.scss";
const Login = () =>
{
    const [loginName, setLoginName] = useState("");
    const [password, setPassword] = useState("");
    
    const [loginNameInvalid, setLoginNameInvalid] = useState(false);
    const [ passwordInvalid, setPasswordInvalid ] = useState( false );
    
    const dispatch = useDispatch();
    const navigate = useNavigate();

    const handleSubmit = () =>
    {
        let empty = false;
        if ( loginName == "" )
        {
            setLoginNameInvalid( "This field cannot be empty." );
            empty = true
        }
        if ( password == "" )
        {
            setPasswordInvalid( "This field cannot be empty." );
            empty = true
        }

        if ( empty )
        {
            return;    
        }

        const clientData = {
            id: loginName,
            password: password,
        }
        login( clientData ).then( result =>
        {
            const user = {
                loginName:result.data.data.id
            }
            function setCookie(name, value, expirationDays) {
                const date = new Date();
                date.setTime(date.getTime() + (expirationDays * 24 * 60 * 60 * 1000));
                const expires = "expires=" + date.toUTCString();
                document.cookie = name + "=" + value + ";" + expires + ";path=/";
            }

            setCookie( "accessToken", result.data.token, 1 );
            setCookie("id", user.loginName, 365)
            dispatch( updateUser( user ) );
            navigate("/")
        } )
        .catch( (result) =>
        {
            if ( result.response.status == "401" )
            {
                setLoginNameInvalid( "Login name or password is incorrect" );
                setPasswordInvalid( "Login name or password is incorrect" );
                return;
            }   
        })
    }

    return (
        <div className="auth">
            <div className="auth-container tag">
                <h2>Login to</h2>
                <form>
                    <Input fieldName="Login name:" type="text" id="login-name" placeholder="example" value={loginName} onChange={e=>setLoginName(e.target.value)} invalidText={loginNameInvalid} removeInvalidStatus={e=>setLoginNameInvalid("")} />
                    <Input fieldName="Mật khẩu:" type="password" id="password" placeholder="• • • • • • • •" value={password} onChange={e=>{setPassword(e.target.value)}} invalidText={passwordInvalid} removeInvalidStatus={e=>setPasswordInvalid("")} />
                <span>Bạn <Link to="/forgot" style={{textDecoration:"underline", fontWeight:"500"}}>quên</Link> mật khẩu?</span>
                </form>
                <button onClick={e=>handleSubmit()}>Đăng nhập</button>
                <div className="auth-container-orText">or</div>
                <Link className="align-center" style={{textDecoration:"underline", fontWeight:"500", padding:"10px"}} to="/register">Đăng ký</Link>
            </div>
        </div>
    )
}

export default Login