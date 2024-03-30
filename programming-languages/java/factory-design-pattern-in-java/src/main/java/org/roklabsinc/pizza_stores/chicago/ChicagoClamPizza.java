package org.roklabsinc.chicago_store;

import org.roklabsinc.Pizza;

public class ChicagoClamPizza implements Pizza {
    @Override
    public void prepare() {
        System.out.println("Preparing Chicago Clam Pizza");
    }

    @Override
    public void bake() {
        System.out.println("Baking Chicago Clam Pizza");
    }

    @Override
    public void cut() {
        System.out.println("Cutting Chicago Clam Pizza");
    }

    @Override
    public void box() {
        System.out.println("Boxing Chicago Clam Pizza");
    }
}
