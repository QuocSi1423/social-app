import React, { useState } from "react";
import BriefUserInformation from "../components/BriefUserInformation";
import { useSelector } from "react-redux/es/hooks/useSelector";
import "../styles/views/Home.scss"

const Home = () =>
{
    const user = useSelector( state => state.user );

    return (
        <div className="home">
            <div className="home-user tag">
                <img className="briefUserInformation-image" src={ user.avatarUrl } alt="" />
                <button className="home-user-posting">{user.userName}, write something</button>
            </div>
        </div>
    )
}

export default Home