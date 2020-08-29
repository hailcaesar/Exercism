#[derive(Debug, PartialEq)]
pub enum Comparison {
    Equal,
    Sublist,
    Superlist,
    Unequal,
}

pub fn sublist<T: PartialEq>(a: &[T], b: &[T]) -> Comparison {
    let sublist = |a: &[T], b: &[T], c| 
        if a.windows(b.len()).any(|x| x == b){ c } else { Comparison::Unequal };

    match (a.len(), b.len()) {
        (0, 0) => Comparison::Equal,
        (_, 0) => Comparison::Superlist,
        (0, _) => Comparison::Sublist,
        (a_size, b_size) if a_size < b_size => sublist(b, a, Comparison::Sublist),
        (b_size, a_size) if b_size > a_size => sublist(a, b, Comparison::Superlist),
        _ => if a == b { Comparison::Equal } else { Comparison::Unequal }
    }
}
