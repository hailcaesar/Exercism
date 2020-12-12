use std::cmp;

pub fn annotate(minefield: &[&str]) -> Vec<String> {
    fn count_mines(row: usize, col: usize, mines: &[&str]) -> char {
        let rstart = row.saturating_sub(1);
        let rend = cmp::min(mines.len() - 1, row + 1);

        let cstart = col.saturating_sub(1);
        let cend = cmp::min(mines[row].len() - 1, col + 1);

        let mut sum = 0;
        for row in mines.iter().take(rend + 1).skip(rstart) {
            let row_bytes = row.as_bytes();
            for ch in row_bytes.iter().take(cend + 1).skip(cstart) {
                if *ch == b'*' {
                    sum += 1;
                }
            }
        }

        if sum == 0 {
            ' '
        } else {
            (sum + b'0') as char
        }
    }

    let mut ans: Vec<String> = vec!["".to_string(); minefield.len()];

    for (row_idx, row) in minefield.iter().enumerate() {
        for (col_idx, ch) in row.chars().enumerate() {
            if ch == '*' {
                ans[row_idx].push(ch)
            } else {
                ans[row_idx].push(count_mines(row_idx, col_idx, minefield));
            }
        }
    }

    ans
}
