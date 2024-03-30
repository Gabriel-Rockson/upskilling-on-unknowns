package org.roklabsinc.pizza_stores.chicago;

import org.roklabsinc.Pizza;
import org.roklabsinc.PizzaStore;
import org.roklabsinc.PizzaType;

public class ChicagoPizzaStore extends PizzaStore {
    @Override
    public Pizza createPizza(PizzaType pizzaType) {
        return switch (pizzaType) {
            case PizzaType.VEGGIE -> new ChicagoVeggiePizza();
            case CHEESE -> new ChicagoCheezePizza();
            case CLAM -> new ChicagoClamPizza();
            case PEPPERONI -> new ChicagoPepperoniPizza();
        };
    }
}
