package org.roklabsinc.pizza_stores.chicago.ny;

import org.roklabsinc.Pizza;

public class NYClamPizza implements Pizza {
    @Override
    public void prepare() {
        System.out.println("Preparing NY Clam Pizza");
    }

    @Override
    public void bake() {
        System.out.println("Baking NY Clam Pizza");
    }

    @Override
    public void cut() {
        System.out.println("Cutting NY Clam Pizza");
    }

    @Override
    public void box() {
        System.out.println("Boxing NY Clam Pizza");
    }
}
