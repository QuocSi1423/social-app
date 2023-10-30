import React from "react";
import {NavLink} from "react-router-dom"
import { BiHomeAlt, BiNotification, BiUserCircle, BiCog } from "react-icons/bi"
import {HiMiniUserGroup} from "react-icons/hi2"
import "../styles/NavigationBar.scss"
const NavigationBar = () =>
{
    const iconSize = 36;
    return (
        <div className="navigationBar">
            <NavLink className="navigationBar-item" to="/"> <BiHomeAlt size={iconSize}/> </NavLink>
            <NavLink className="navigationBar-item" to="/me"> <BiUserCircle size={iconSize}/> </NavLink>
            <NavLink className="navigationBar-item" to="/notification"> <BiNotification size={ iconSize } /> </NavLink>
            <NavLink className="navigationBar-item" to="/group"> <HiMiniUserGroup size={ iconSize } /> </NavLink>
            <NavLink className="navigationBar-item" to="setting" > <BiCog size={iconSize}/> </NavLink>
        </div>
    )
}

export default NavigationBar