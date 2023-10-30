import React from "react";   
import "../styles/Component.scss"
import { Link } from "react-router-dom";

const BriefUserInformation = (props) =>
{
    const { avatarUrl, username } = props;
    return (
        <div className="briefUserInformation">
            <Link>
                <img className="briefUserInformation-image" src={ avatarUrl } alt="" />
                <span className="briefUserInformation-name">{username}</span>
            </Link>
        </div>
    )
}

export default BriefUserInformation;