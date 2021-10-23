import React from 'react';
import Item from './Item.js'
class Home extends React.Component {
    
    constructor(props) {
        super(props);
        this.state = {
            imgArray: []
        };
        this.selectItem = this.selectItem.bind(this);

    }

    componentDidMount() {
        // TODO: Make API call to get products
        this.setState({
            imgArray: [{name: "Pizza", price: 5.99, link:"https://upload.wikimedia.org/wikipedia/commons/a/a3/Eq_it-na_pizza-margherita_sep2005_sml.jpg"},
            {name:"Ice cream", price:2.99, link: "https://upload.wikimedia.org/wikipedia/commons/thumb/2/2e/Ice_cream_with_whipped_cream%2C_chocolate_syrup%2C_and_a_wafer_%28cropped%29.jpg/1024px-Ice_cream_with_whipped_cream%2C_chocolate_syrup%2C_and_a_wafer_%28cropped%29.jpg"}]
        });
    }


    render(){
        let imageArrayRender = this.state.imgArray.map(image => {
            return (
                <Item 
                    name={image.name}
                    price={image.price}
                    selectItem={this.selectItem}
                    link={image.link}
                />
            );
        });
        return (
            <>
                {imageArrayRender}
            </>
        );
    }

    selectItem = (itemName) => {
        console.log(itemName);
    }
}


export default Home;