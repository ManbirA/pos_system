import React from 'react';
import './Item.css';

class Item extends React.Component {

    render() {
        const {name, price, selectItem, link} = this.props;
        return (
            <div className="product">
                <img className="product-picture" src={require(`../pictures/${link}`).default} alt="Logo" onClick={() => selectItem("React Logo click!!")}/>
                <h6>{name}</h6>
                <h6>{price}</h6>
            </div>
        );
    }
}


export default Item;