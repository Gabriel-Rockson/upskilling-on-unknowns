package org.roklabsinc;

public class Main {
    public static void main(String[] args) {
        SimplePizzaFactory simplePizzaFactory = new SimplePizzaFactory();
        PizzaStore store = new PizzaStore(simplePizzaFactory);

        Pizza pizza = store.orderPizza(PizzaType.VEGGIE);

        System.out.println(pizza);
    }
}
