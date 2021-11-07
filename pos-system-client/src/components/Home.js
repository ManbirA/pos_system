import React from 'react';
import Item from './Item.js'


class Home extends React.Component {
    
    constructor(props) {
        super(props);
        this.state = {
            propductArray: []
        };
        this.selectItem = this.selectItem.bind(this);

    }

    componentDidMount() {
        fetch("http://localhost:8080/products")
            .then(products => products.json())
            .then((result) => {
                var allProduct = [];
                for(var key of Object.keys(result)){
                    allProduct.push({name: result[key].name, price: result[key].price, inventory: result[key].inventory });
                }
                this.setState({
                    propductArray: allProduct
                });
            }, (error) => {
                console.log(error);
            });
    }


    render(){
        let productArrayRender = this.state.propductArray.map(product => {
            return (
                <Item 
                    name={product.name}
                    price={product.price}
                    selectItem={this.selectItem}
                    link={"logo512.png"}

                />
            );
        });
        return (
            <>
                {productArrayRender}
            </>
        );
    }

    selectItem = (itemName) => {
        console.log(itemName);
    }
}


export default Home;