package main

type Pedido struct {
	Id         int64  `json:"id"`
	Producto   string `json:"producto"`
	Cateegoria string `json:"categoria"`
	Cantidad   int64  `json:"cantidad"`
}
