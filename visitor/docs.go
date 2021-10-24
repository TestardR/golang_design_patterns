package visitor

/*
1. Propagate an Accept(v *Visitor) method throughout the entire hirarchy
2. Create a visitor with VisitFoor(f Foo), VisitBar(b Bar), for each element in the hierarchy
3. Each Accept() simply calls Visitor.VisitXxx(self)
*/
