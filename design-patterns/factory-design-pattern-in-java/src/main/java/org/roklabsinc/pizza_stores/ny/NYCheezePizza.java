package org.roklabsinc.pizza_stores.chicago.ny;

import org.roklabsinc.Pizza;

public class NYCheezePizza implements Pizza {
    @Override
    public void prepare() {
        System.out.println("Preparing NY Cheeze Pizza");
    }

    @Override
    public void bake() {
        System.out.println("Baking NY Cheeze Pizza");
    }

    @Override
    public void cut() {
        System.out.println("Cutting NY Cheeze Pizza");
    }

    @Override
    public void box() {
        System.out.println("Boxing NY Cheeze Pizza");
    }
}
