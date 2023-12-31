import axios from "axios"

export const register = (registerData) =>
{
    return axios.post( "http://localhost:5173/v1/register", registerData );
}

export const login = ( loginData ) =>
{
    return axios.post( "http://localhost:5173/v1/login", loginData );
}