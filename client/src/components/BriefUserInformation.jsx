import React from "react";   
import "../styles/Component.scss"
import { Link } from "react-router-dom";

const BriefUserInformation = (props) =>
{
    
    const { avatarUrl, username, avatarSize, nameSize } = props;
    return (
        <div className="briefUserInformation">
            <Link>
                <img style={ avatarSize ? { width: `${ avatarSize }px`,height:`${avatarSize}px`}:{}} className="briefUserInformation-image" src={ avatarUrl } alt="" />
                <span style={ nameSize ? {fontSize:`${nameSize}px`}:{}} className="briefUserInformation-name">{username}</span>
            </Link>
        </div>
    )
}

export default BriefUserInformation;