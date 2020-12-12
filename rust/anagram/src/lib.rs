use std::collections::HashMap;
use std::collections::HashSet;

fn count_letters(word: &str) -> HashMap<char, u32> {
    let mut count: HashMap<char, u32> = HashMap::new();
    let word_lower = word.to_lowercase();
    for ch in word_lower.chars() {
        *count.entry(ch).or_insert(0) += 1;
    }

    count
}

pub fn anagrams_for<'a>(word: &str, possible_anagrams: &'a [&str]) -> HashSet<&'a str> {
    let word = word.to_lowercase();
    let required = count_letters(&word);
    let mut ans = HashSet::<&str>::new();

    for candidate in possible_anagrams {
        let count = count_letters(candidate);

        if count == required && candidate.to_lowercase() != word {
            ans.insert(candidate);
        }
    }

    ans
}
