import axios from "axios";

export const getBriefUserInformation = (userId) =>
{
    return axios.get(`http://localhost:5173/v1/users/${userId}/informations/brief`)
}