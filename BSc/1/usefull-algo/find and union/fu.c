#include <bits/stdc++.h>
#define T int
using namespace std;

// Drzewo find and union
struct Set {
    T el;
    Set* up;
    unsigned int rank;
    vector<Set*> next;

    // konstruktor przyjmuje wartoœæ
    Set(T elem) {
        this->el = elem;
        this->rank = 0;
        this->next = vector<Set*>();
        this->up = this;
    }

    // idzie do korzenie drzewa
    Set* goUp() {
        if(this->up != this) {
            return this->up->goUp();
        }
        return this;
    }

    // zwraca wartoœc w korzeniu
    T find() {
        return this.goUp()->el;
    }

    // sprzawdza czy dwa drzewa s¹ tym samym
    bool equivalent(Set s) {
        return s.goUp() == this->goUp();
    }

    // oprecja union
    Set* un(Set s) {
        Set *root = this->goUp();
        Set *sRoot = s.goUp();
        if(root != sRoot) {
            Set *newRoot, *other;
            if(root->rank > sRoot->rank) {
                newRoot = root;
                other = sRoot;
            } else {
                newRoot = sRoot;
                other = root;
                if(root->rank == sRoot->rank) {
                    sRoot->rank = root->rank +1;
                }
            }
            other->up = newRoot;
            newRoot->next.push_back(other);
            return newRoot;
        }
        return this;
    }

    // zwrócenie wszystkich elementów w drzewie
    vector<T> elements() {
        return this->goUp()->traverse(vector<T>());
    }

    // funckja pomocnicza do zwracania elementów
    vector<T> traverse(vector<T> els) {
        els.push_back(this->el);
        for(Set *child : this->next) {
            child->traverse(els);
        }
        return els;
    }
};


//template<typename T>
//class FindUnion {
//  public:
//    FindUnion() {
//        this.setCounter = 0;
//    }
//    void makeSet(T el) {
//        this.Set(el);
//        this.setCounter++;
//    }
//    int noSets() {
//        return this.setCounter;
//    }
//    void union(Set* a, Set *b) {
//
//    }
//  private:
//    int setCounter;
//
//};
