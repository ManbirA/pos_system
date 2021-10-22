import React from 'react';
import picture from '../pictures/logo512.png'
class Item extends React.Component {

    // constructor(props) {
    //     super(props);

    // }
    render() {
        return (
            <>
                <img src={picture} alt="Logo" onClick={() => this.props.selectItem("React Logo click!!")}/>
            </>
        );
    }
}


export default Item;