export interface Stock {
    ticker: string;
    company: string;
    brokerage: string;
    action: string;
    rating_from: string;
    rating_to: string;
    target_to: number;
    target_from: number;
    score: number;
}