package main

func createPedido(Pedido Pedido) error {
	bd, err := getDB()
	if err != nil {
		return err
	}
	_, err = bd.Exec("INSERT INTO pedido (producto, categoria, cantidad) VALUES (?, ?, ?)", Pedido.Producto, Pedido.Cateegoria, Pedido.Cantidad)
	return err
}

func deletePedido(id int64) error {

	bd, err := getDB()
	if err != nil {
		return err
	}
	_, err = bd.Exec("DELETE FROM pedido WHERE id = ?", id)
	return err
}

// It takes the ID to make the update
func updatePedido(Pedido Pedido) error {
	bd, err := getDB()
	if err != nil {
		return err
	}
	_, err = bd.Exec("UPDATE pedido SET producto = ?, categoria = ?, cantidad = ? WHERE id = ?", Pedido.Producto, Pedido.Cateegoria, Pedido.Cantidad, Pedido.Id)
	return err
}
func getPedidos() ([]Pedido, error) {
	//Declare an array because if there's error, we return it empty
	Pedidos := []Pedido{}
	bd, err := getDB()
	if err != nil {
		return Pedidos, err
	}
	// Get rows so we can iterate them
	rows, err := bd.Query("SELECT id, producto, categoria, cantidad FROM pedido")
	if err != nil {
		return Pedidos, err
	}
	// Iterate rows...
	for rows.Next() {
		// In each step, scan one row
		var Pedido Pedido
		err = rows.Scan(&Pedido.Id, &Pedido.Producto, &Pedido.Cateegoria, &Pedido.Cantidad)
		if err != nil {
			return Pedidos, err
		}
		// and append it to the array
		Pedidos = append(Pedidos, Pedido)
	}
	return Pedidos, nil
}

func getPedidoById(id int64) (Pedido, error) {
	var Pedido Pedido
	bd, err := getDB()
	if err != nil {
		return Pedido, err
	}
	row := bd.QueryRow("SELECT id, producto, categoria, cantidad FROM pedido WHERE id = ?", id)
	err = row.Scan(&Pedido.Id, &Pedido.Producto, &Pedido.Cateegoria, &Pedido.Cantidad)
	if err != nil {
		return Pedido, err
	}
	// Success!
	return Pedido, nil
}
