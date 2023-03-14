use std::collections::HashMap;

impl Solution {

    pub fn counter(s: String) -> HashMap<char, i32> {
        let mut map: HashMap<char, i32> = HashMap::new();
        for c in s.chars() {
            let x = map.entry(c).or_insert(0);
            *x += 1
        }
        return map;
    }

    pub fn is_anagram(s: String, t: String) -> bool {
        if (s.len() != t.len()) {
            return false;
        }

        let mut s_map: HashMap<char, i32> = Solution::counter(s);
        let mut t_map: HashMap<char, i32> = Solution::counter(t);
        
        
    }
}
