package db

import (
	"gopkg.in/mgo.v2/bson"
	"products.akash.com/log"
	"products.akash.com/model"
)

func FindOneWithId(id string) *model.Product {

	c := Session.DB("shopping_cart_dev_products").C("product-collection")

	data := &model.Product{}
	err := c.Find(bson.M{"id": id}).One(data)
	if err != nil {
		log.Error("Error while finding product with id ", id)
		return nil
	}

	log.Info("Found product for this id : ", *data)

	return data
}

func FindAll() []model.Product {

	c := Session.DB("shopping_cart_dev_products").C("product-collection")

	log.Info("Going to find all products")
	var results []model.Product
	err := c.Find(nil).All(&results)
	if err != nil {
		log.Error("Error while querying all products ", err)
		panic(err)
	} else {
		log.Info("Found products : ", results)
	}
	return results
}

func SaveProduct(product model.Product) {

	c := Session.DB("shopping_cart_dev_products").C("product-collection")

	log.Info("Trying to save product : ", product)

	if err := c.Insert(product); err != nil {
		log.Error("Error while saving product ", err)
		panic(err)
	}

	log.Info("Product added ")
}

func DeleteProduct(id string) bool {

	c := Session.DB("shopping_cart_dev_products").C("product-collection")

	log.Info("Trying to delete product with id : ", id)

	err := c.Remove(bson.M{"id": id})
	if err != nil {
		log.Error("Error while deleting product with id ", err)
		return false
	} else {
		log.Info("Product deleted")
		return true
	}
}
