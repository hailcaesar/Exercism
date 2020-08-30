pub fn find(array: &[i32], key: i32) -> Option<usize> {
    let mut l = 0;
    let mut h = array.len() as i32 - 1;

    while l <= h {
        let m = l + (h - l) / 2;
        let m_idx: usize = m as usize;

        if array[m_idx] == key {
            return Some(m_idx);
        }
        else if key > array[m_idx] {
            l = m + 1;
        }
        else{
            h = m - 1;
        }
    }
    return None
}
    
