
use std::collections::HashMap;

fn count_letters(word: &str) -> HashMap<char, i32> {
    word.to_lowercase()
        .chars()
        .fold(HashMap::new(), |mut counts, c| {
            *counts.entry(c).or_insert(0) += 1;
            counts
        })
}

pub fn anagrams_for<'a>(word: &str, candidates: &[&'a str]) -> Vec<&'a str> {
    let word       = word.to_lowercase();
    let wordcounts = count_letters(&word);
    candidates.iter()
              .filter(|w| w.to_lowercase() != word)
              .filter(|w| count_letters(w) == wordcounts)
              .cloned()
              .collect()
}
