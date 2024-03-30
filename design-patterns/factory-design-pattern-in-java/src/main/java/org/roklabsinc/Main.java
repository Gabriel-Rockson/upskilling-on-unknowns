package org.roklabsinc;

import org.roklabsinc.pizza_stores.chicago.ChicagoPizzaStore;
import org.roklabsinc.pizza_stores.chicago.ny.NYPizzaStore;

public class Main {
    public static void main(String[] args) {

        // Nathan's order
        PizzaStore nyPizzaStore = new NYPizzaStore();
        nyPizzaStore.orderPizza(PizzaType.CLAM);

        // Ethan's order
        PizzaStore chicagoPizzaStore = new ChicagoPizzaStore();
        chicagoPizzaStore.orderPizza(PizzaType.PEPPERONI);
    }
}
