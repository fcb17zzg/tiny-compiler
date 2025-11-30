package compiler

import "testing"

func TestDefine(t *testing.T) {
	expected := map[string]Symbol{
		"a": {Name: "a", Scope: GlobalScope, Index: 0},
		"b": {Name: "b", Scope: GlobalScope, Index: 1},
		"c": {Name: "c", Scope: LocalScope, Index: 0},
		"d": {Name: "d", Scope: LocalScope, Index: 1},
		"e": {Name: "e", Scope: LocalScope, Index: 0},
		"f": {Name: "f", Scope: LocalScope, Index: 1},
	}

	global := NewSymbolTable()

	a := global.Define("a")
	if a != expected["a"] {
		t.Errorf("expected=%+v, got=%+v", expected["a"], a)
	}

	b := global.Define("b")
	if b != expected["b"] {
		t.Errorf("expected=%+v, got=%+v", expected["b"], b)
	}

	firstLocal := NewEnclosedSymbolTable(global)
	c := firstLocal.Define("c")
	if c != expected["c"] {
		t.Errorf("expected=%+v, got=%+v", expected["c"], c)
	}

	d := firstLocal.Define("d")
	if d != expected["d"] {
		t.Errorf("expected=%+v, got=%+v", expected["d"], d)
	}

	secondLocal := NewEnclosedSymbolTable(firstLocal)

	e := secondLocal.Define("e")
	if e != expected["e"] {
		t.Errorf("expected=%+v, got=%+v", expected["e"], e)
	}

	f := secondLocal.Define("f")
	if f != expected["f"] {
		t.Errorf("expected=%+v, got=%+v", expected["f"], f)
	}
}

func TestResolveGlobal(t *testing.T) {
	global := NewSymbolTable()
	global.Define("a")
	global.Define("b")

	expected := []Symbol{
		Symbol{Name: "a", Scope: GlobalScope, Index: 0},
		Symbol{Name: "b", Scope: GlobalScope, Index: 1},
	}

	for _, sym := range expected {
		result, ok := global.Resolve(sym.Name)
		if !ok {
			t.Fatalf("failed to resolve %s", sym.Name)
		}

		if result != sym {
			t.Errorf("expected=%+v, got=%+v", sym, result)
		}
	}
}

func TestResolveLocal(t *testing.T) {
	global := NewSymbolTable()
	global.Define("a")
	global.Define("b")

	local := NewEnclosedSymbolTable(global)
	local.Define("c")
	local.Define("d")

	expected := []Symbol{
		{Name: "a", Scope: GlobalScope, Index: 0},
		{Name: "b", Scope: GlobalScope, Index: 1},
		{Name: "c", Scope: LocalScope, Index: 0},
		{Name: "d", Scope: LocalScope, Index: 1},
	}

	for _, sym := range expected {
		resolved, ok := local.Resolve(sym.Name)
		if !ok {
			t.Fatalf("symbol %s not resolved", sym.Name)
			continue
		}
		if resolved != sym {
			t.Errorf("symbol %s resolved incorrectly. got=%+v, want=%+v",
				sym.Name, resolved, sym)
		}
	}
}

func TestResolveNestedLocal(t *testing.T) {
	global := NewSymbolTable()
	global.Define("a")
	global.Define("b")

	local1 := NewEnclosedSymbolTable(global)
	local1.Define("c")
	local1.Define("d")

	local2 := NewEnclosedSymbolTable(local1)
	local2.Define("e")
	local2.Define("f")

	tests := []struct {
		table           *SymbolTable
		expectedSymbols []Symbol
	}{
		{
			local1,
			[]Symbol{
				Symbol{Name: "a", Scope: GlobalScope, Index: 0},
				Symbol{Name: "b", Scope: GlobalScope, Index: 1},
				Symbol{Name: "c", Scope: LocalScope, Index: 0},
				Symbol{Name: "d", Scope: LocalScope, Index: 1},
			},
		},
		{
			local2,
			[]Symbol{
				Symbol{Name: "a", Scope: GlobalScope, Index: 0},
				Symbol{Name: "b", Scope: GlobalScope, Index: 1},
				Symbol{Name: "e", Scope: LocalScope, Index: 0},
				Symbol{Name: "f", Scope: LocalScope, Index: 1},
			},
		},
	}

	for _, tt := range tests {
		for _, sym := range tt.expectedSymbols {
			resolved, ok := tt.table.Resolve(sym.Name)
			if !ok {
				t.Fatalf("symbol %s not resolved", sym.Name)
				continue
			}
			if resolved != sym {
				t.Errorf("symbol %s resolved incorrectly. got=%+v, want=%+v",
					sym.Name, resolved, sym)
			}
		}
	}
}
